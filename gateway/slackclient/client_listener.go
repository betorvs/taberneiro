package slackclient

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/betorvs/taberneiro/config"
	"github.com/nlopes/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

// Slack struct for our slackbot
type Slack struct {
	Name  string
	Token string

	User   string
	UserID string

	Client *slack.Client
}

// New returns a new instance of the Slack struct, primary for our slackbot
func New() (*Slack, error) {
	return &Slack{Client: slack.New(config.SlackToken), Token: config.SlackToken, Name: config.SlackAppName}, nil
}

// Run func
func (s *Slack) Run(ctx context.Context) error {
	authTest, err := s.Client.AuthTest()
	if err != nil {
		return fmt.Errorf("did not authenticate: %+v", err)
	}

	s.User = authTest.User
	s.UserID = authTest.UserID

	log.Printf("[INFO] bot is now registered as %v (%v)\n", s.User, s.UserID)

	go s.run(ctx)
	return nil
}

// run func
func (s *Slack) run(ctx context.Context) {

	rtm := s.Client.NewRTM()
	go rtm.ManageConnection()

	fmt.Println("[INFO] now listening for incoming messages...")
	for msg := range rtm.IncomingEvents {
		switch ev := msg.Data.(type) {

		case *slack.ConnectedEvent:
			log.Println("Infos:", ev.Info)
			log.Println("Connection counter:", ev.ConnectionCount)
			attachment := slack.Attachment{
				Text:     "My Menu",
				ImageURL: config.MenuImageLink,
			}
			rtm.SendMessage(rtm.NewOutgoingMessage(config.HelloMessage, config.ChannelID))
			go SimpleMessage(config.ChannelID, attachment)

		case *slack.MessageEvent:
			// fmt.Printf("Message: %v\n", ev)
			if err := s.messageHandler(ev); err != nil {
				log.Printf("[ERROR] Failed to handle message: %s", err)
			}

		case *slack.RTMError:
			log.Printf("Error: %s\n", ev.Error())

		case *slack.InvalidAuthEvent:
			log.Printf("Invalid credentials")
			return

		default:

			// Ignore other events..
			// log.Printf("Unexpected: %v\n", msg.Data)
		}
	}

}

// messageHandler func
func (s *Slack) messageHandler(ev *slack.MessageEvent) error {
	// Only response in specific channel. Ignore else.
	if ev.Channel != config.ChannelID {
		log.Printf("%v %v", ev.Channel, ev.Msg.Text)
		return nil
	}
	// parse subtype message to exclude false errors
	if ev.Msg.SubType == "message_changed" || ev.Msg.SubType == "message_deleted" || ev.Msg.SubType == "bot_message" {
		return nil
	}
	// Parse message
	menuText := fmt.Sprintf("%s", config.MenuText)
	text := fmt.Sprintf("%s : %s", config.Suggestion, config.MagicWord)
	m := strings.Split(strings.TrimSpace(ev.Msg.Text), " ")[1:]
	if len(m) == 0 || m[0] != config.MagicWord {
		attachment := slack.Attachment{
			Text: text,
		}
		if strings.Contains(ev.Msg.Text, "@"+s.UserID) {
			go SimpleMessage(ev.Channel, attachment)
		}
		return nil
	}
	options := make([]slack.AttachmentActionOption, 0)
	err := json.Unmarshal([]byte(config.SelectMenu), &options)
	if err != nil {
		log.Printf("Error reading config.SelectMenu : %s", err)
	}

	attachment := slack.Attachment{
		Text:       menuText,
		CallbackID: "order",
		Actions: []slack.AttachmentAction{
			{
				Name:    actionSelect,
				Type:    "select",
				Options: options,
			},
			{
				Name:  actionCancel,
				Text:  "Cancel",
				Type:  "button",
				Style: "danger",
			},
		},
	}
	go SimpleMessage(ev.Channel, attachment)
	return nil
}

// SimpleMessage func
func SimpleMessage(channel string, attachment slack.Attachment) error {
	s, err := New()
	if err != nil {
		log.Printf("Error creating slack client: %s", err)
	}
	if _, _, err := s.Client.PostMessage(channel, slack.MsgOptionAttachments(attachment)); err != nil {
		return fmt.Errorf("failed to post message: %v", err)
	}
	return nil
}

// EphemeralMessage func
func EphemeralMessage(channel string, user string, attachment slack.Attachment) error {
	s, err := New()
	if err != nil {
		log.Printf("Error creating slack client: %s", err)
	}
	if _, err := s.Client.PostEphemeral(channel, user, slack.MsgOptionAttachments(attachment)); err != nil {
		return fmt.Errorf("failed to post message: %v", err)
	}
	return nil
}

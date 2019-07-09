package usecase

import (
	"fmt"
	"log"
	"strings"

	"github.com/betorvs/taberneiro/config"
	"github.com/betorvs/taberneiro/gateway/slackclient"
	"github.com/labstack/echo"
	"github.com/nlopes/slack"
)

const (
	// action is used for slack attament action.
	actionSelect = "select"
	actionStart  = "start"
	actionCancel = "cancel"
)

// ActionEvent func
func ActionEvent(data *slack.InteractionCallback, c echo.Context) (slack.Message, error) {
	var res slack.Message
	option := data.ActionCallback
	switch option.AttachmentActions[0].Name {
	case actionSelect:
		username := data.User.Name
		value := option.AttachmentActions[0].SelectedOptions[0].Value
		originalMessage := data.OriginalMessage
		originalMessage.ReplaceOriginal = true
		originalMessage.Attachments[0].Text = fmt.Sprintf("%s OK to order %s", username, strings.Title(value))
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{
			{
				Name:  actionStart,
				Text:  "Yes",
				Type:  "button",
				Value: "start",
				Style: "primary",
			},
			{
				Name:  actionCancel,
				Text:  "No",
				Type:  "button",
				Style: "danger",
			},
		}
		res = originalMessage
	case actionStart:
		title := ":ok: i accepted that!"
		value := ""
		originalMessage := data.OriginalMessage
		originalMessage.ReplaceOriginal = true
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
		originalMessage.Attachments[0].Fields = []slack.AttachmentField{
			{
				Title: title,
				Value: value,
				Short: false,
			},
		}
		text := originalMessage.Attachments[0].Text
		text = strings.Replace(text, "OK to order", "ordered", -1)
		attachment := slack.Attachment{
			Text: fmt.Sprintf("Order Submitted: %s", text),
		}
		go slackclient.SimpleMessage(config.DirectMessage, attachment)
		res = originalMessage

	case actionCancel:
		title := fmt.Sprintf(":x: @%s canceled the request", data.User.Name)
		value := ""
		originalMessage := data.OriginalMessage
		originalMessage.ReplaceOriginal = true
		originalMessage.Attachments[0].Actions = []slack.AttachmentAction{} // empty buttons
		originalMessage.Attachments[0].Fields = []slack.AttachmentField{
			{
				Title: title,
				Value: value,
				Short: false,
			},
		}
		res = originalMessage
	default:
		log.Printf("[ERROR] Invalid action was submitted: %s", option.AttachmentActions[0].Name)
	}
	return res, nil
}

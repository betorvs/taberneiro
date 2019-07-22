package usecase

import (
	"github.com/betorvs/taberneiro/config"
	"github.com/betorvs/taberneiro/domain"
	"github.com/betorvs/taberneiro/gateway/slackclient"
	"github.com/nlopes/slack"
)

// PostMessage func
func PostMessage(data *domain.Message) (string, error) {
	text := data.Text
	attachment := slack.Attachment{
		Text: text,
	}
	go slackclient.SimpleMessage(config.ChannelID, attachment)
	return "OK", nil
}

// ValidateHeader func
func ValidateHeader(header string) bool {
	if config.HeaderMagic != header {
		return false
	}
	return true

}

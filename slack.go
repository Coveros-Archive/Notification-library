package notification_library

import (
	"encoding/json"
	"fmt"
)

const (
	SlackAPIUrl = "https://slack.com/api/chat.postMessage"
)

type slackClient struct {
	SlackApiUrl    string
	BotAccessToken string
}

type slackData struct {
	Channel     string `json:"channel"`
	Text        string `json:"text"`
	Attachments []slackAttachment `json:"attachments"`
}

type slackAttachment struct {
	MrkdwnIn   []string `json:"mrkdwn_in"`
	Color      string   `json:"color"`
	Pretext    string   `json:"pretext"`
	AuthorName string   `json:"author_name"`
	AuthorLink string   `json:"author_link"`
	AuthorIcon string   `json:"author_icon"`
	Title      string   `json:"title"`
	TitleLink  string   `json:"title_link"`
	Text       string   `json:"text"`
	SlackFields []slackField `json:"fields"`
}

type slackField struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"short"`
}

func newSlack(botAccessToken string) *slackClient {
	return &slackClient{
		SlackApiUrl:    SlackAPIUrl,
		BotAccessToken: botAccessToken,
	}
}

func (s *slackClient) SendMsg(n NotifyTemplate) error {
	slackData := notifyTemplateToSlackTemplate(n)
	jsonByte , errMarshalling := json.Marshal(slackData)
	if errMarshalling != nil {
		return errMarshalling
	}
	headers := map[string]string{"Authorization": fmt.Sprintf("Bearer %v", s.BotAccessToken), "Content-type": "application/json"}
	return httpPost(s.SlackApiUrl, headers, jsonByte)
}


func getColorBasedOnEventType (eventType NotifyEventType) string {
	var color string
	switch eventType {
	case Success:
		color = "good"
	case Failure:
		color = "danger"
	case Warning:
		color = "warning"
	default:
		color = "#00FFFF" // cyan color
	}
	return color
}

func notifyTemplateFieldsToSlackFields(n NotifyTemplate) []slackField {
	var sFields []slackField
	for k, v := range n.Fields {
		sField := slackField{
			Title: k,
			Value: v,
		}
		sFields = append(sFields, sField)
	}
	return sFields
}

func notifyTemplateToSlackTemplate(n NotifyTemplate) slackData {
	sData := slackData{
		Channel: n.Channel,
	}
	sAttachments := []slackAttachment{
		{
			Title: n.Title,
			Text: n.Text,
			Color: getColorBasedOnEventType(n.EventType),
			SlackFields: notifyTemplateFieldsToSlackFields(n),
		},
	}
	sData.Attachments  = sAttachments
	return sData
}


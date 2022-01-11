package helper

import (
	"bekci/configmap"
	"bekci/model"
	"github.com/ashwanthkumar/slack-go-webhook"
	log "github.com/sirupsen/logrus"
)

func Slack(msg, env string, cluster model.Clusters) {
	webhookUrl := configmap.Config.WEBHOOK_URL
	url := "https://docs.google.com/spreadsheets/d/" + configmap.Config.SHEET_ID + "/edit?ts=5f524b59#gid=0"
	if env == "TEST" {
		url = "https://docs.google.com/spreadsheets/d/" + configmap.Config.SHEET_ID + "/edit?ts=5f524b59#gid=280977686"
	}
	attachment1 := slack.Attachment{}
	if cluster.IP == "" {
		log.Warning("Can not detected any errors.")
		attachment1.AddAction(slack.Action{Type: "button", Text: "Dokümanda Kontrol Edin", Url: url, Style: "primary"})
	} else {
		attachment1.AddField(slack.Field{Title: cluster.Hostname, Value: cluster.IP}).AddField(slack.Field{Title: "Expire Date", Value: cluster.Date}).AddField(slack.Field{Title: "Version", Value: cluster.Version})
		attachment1.AddAction(slack.Action{Type: "button", Text: "Dokümanda Kontrol Edin", Url: url, Style: "danger"})
	}
	payload := slack.Payload{
		Text:        msg,
		Username:    "Bekçi",
		Channel:     configmap.Config.CHANNEL,
		IconEmoji:   ":kubernetes:",
		Attachments: []slack.Attachment{attachment1},
	}
	err := slack.Send(webhookUrl, "", payload)
	if len(err) > 0 {
		log.Error(err)
	}
}

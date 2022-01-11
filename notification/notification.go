package notification

import (
	"bekci/helper"
	"bekci/model"
	"strconv"
	"strings"
	"time"
)

func SendNotification(values []model.Clusters, env string) {
	var clusterSize = 0
	var errorCount = 0
	clusterSize += len(values)
	currentTime := time.Now().AddDate(0, 1, 0)
	for _, s := range values {
		dates, _ := time.Parse("02.01.2006", s.Date)
		if strings.Contains(s.Date, "Error") || strings.Contains(s.Version, "Error") {
			helper.Slack("Hata! "+env+" Kubernetes clusterından veri alınırken bir hata oluştu!", env, s)
			errorCount ++
		} else if currentTime.Unix() >= dates.Unix() {
			helper.Slack("Aşağıdaki "+env+" Kubernetes clusterının sertifika bitiş tarihi 1 aydan az kaldı!", env, s)
			errorCount ++
		}
	}
	if errorCount == 0{
		helper.Slack(strconv.Itoa(clusterSize) + " Adet "+env+" Kubernetes clusterının version ve sürüm listesi herhangi bir sorun tespit edilmeden güncellendi. (" + time.Now().Format("02.01.2006 - 15:04") + ")",env, model.Clusters{})
	}
}
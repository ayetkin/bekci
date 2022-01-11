/*
Copyright © 2021 Ali Yetkin info@aliyetkin.com

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package main

import (
	"bekci/notification"
	"bekci/sheetsapi"
	"bekci/update"
	"os"

	runtime "github.com/banzaicloud/logrus-runtime-formatter"
	log "github.com/sirupsen/logrus"
)

func init() {
	formatter := runtime.Formatter{ChildFormatter: &log.TextFormatter{FullTimestamp: true}}
	formatter.Line = false
	formatter.Package = false
	formatter.File = false
	log.SetFormatter(&formatter)
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}

func main() {
	// Kubernetes EUROPA
	VersionsAv, ExpireDatesAv, errAv := update.K8sRowUpdate("", "B3:B100", "A3:A100")
	if errAv != nil {
		log.Error("Yazma işlemine geçilemedi. Veri okunurken bir hata oluştu. (Avrupa)")
	} else {
		sheetsapi.WriteRows("C3", VersionsAv)
		sheetsapi.WriteRows("D3", ExpireDatesAv)
	}

	//Kubernetes ASIA
	VersionsAs, ExpireDatesAs, errAs := update.K8sRowUpdate("", "G3:G100", "F3:F100")
	if errAs != nil {
		log.Error("Yazma işlemine geçilemedi. Veri okunurken bir hata oluştu. (Asya)")
	} else {
		sheetsapi.WriteRows("H3", VersionsAs)
		sheetsapi.WriteRows("I3", ExpireDatesAs)
	}

	// Send notification
	sheetsapi.WriteRows("L2", update.TimeLog())
	notification.SendNotification(update.Values, "PROD")
	update.Values = nil

	//Kubernetes TEST&QA
	VersionsTest, ExpireDatesTest, errTest := update.K8sRowUpdate("Kubernetes Test Envanter List!", "B3:B100", "A3:A100")
	if errTest != nil {
		log.Error("Yazma işlemine geçilemedi. Veri okunurken bir hata oluştu. (Test)")
	} else {
		sheetsapi.WriteRows("Kubernetes Test Envanter List!C3", VersionsTest)
		sheetsapi.WriteRows("Kubernetes Test Envanter List!D3", ExpireDatesTest)
	}
	//Kubernetes SIT&PREPROD
	VersionsSit, ExpireDatesSit, errSit := update.K8sRowUpdate("Kubernetes Test Envanter List!", "G3:G100", "F3:F100")
	if errSit != nil {
		log.Error("Yazma işlemine geçilemedi. Veri okunurken bir hata oluştu. (Sit&PreProd)")
	} else {
		sheetsapi.WriteRows("Kubernetes Test Envanter List!H3", VersionsSit)
		sheetsapi.WriteRows("Kubernetes Test Envanter List!I3", ExpireDatesSit)
	}

	// Send notification
	sheetsapi.WriteRows("Kubernetes Test Envanter List!L2", update.TimeLog())
	notification.SendNotification(update.Values, "TEST")
	update.Values = nil

}

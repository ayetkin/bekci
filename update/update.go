package update

import (
	"bekci/execute"
	"bekci/model"
	"bekci/sheetsapi"
	log "github.com/sirupsen/logrus"
	"time"
)

func TimeLog() []interface{} {
	log.Info("Running timeLog task...")
	time := time.Now()
	var currentTime []interface{}
	currentTime = append(currentTime, time.Format("02.01.2006 15:04:05 (Updated from Bekçi)"))
	return currentTime
}

var Values []model.Clusters

func K8sRowUpdate(sheet, rowsIP, rowsHosts string) ([]interface{}, []interface{}, error) {
	log.Warning("New Google sheet update task was started.")

	IP, err := sheetsapi.GetRows(sheet + rowsIP)
	if err != nil {
		return nil, nil, err
	}

	Hosts, err := sheetsapi.GetRows(sheet + rowsHosts)
	if err != nil {
		return nil, nil, err
	}

	Versions, err := execute.GetK8sVersion(IP)
	if err != nil {
		return nil, nil, err
	}

	ExpireDates, err := execute.GetK8sCertStatus(IP)
	if err != nil {
		return nil, nil, err
	}

	for i, _ := range Hosts {
		Values = append(Values, model.Clusters{
			Hostname: Hosts[i].(string),
			IP:       IP[i].(string),
			Version:  Versions[i].(string),
			Date:     ExpireDates[i].(string),
		})
	}
	return Versions, ExpireDates, nil
}

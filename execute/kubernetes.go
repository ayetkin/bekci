package execute

import (
	"bekci/helper"
	"bekci/model"
	log "github.com/sirupsen/logrus"
)

func GetK8sVersion(rows []interface{}) ([]interface{}, error) {
	var all []interface{}
	for _, s := range rows {
		k8sVersion := model.K8sVersion{}
		err := helper.HttpGet("https://" + s.(string) + ":6443/version", &k8sVersion)
		if err != nil {
			log.Error("Kubernetes api serverdan veri alınırken hata oluştu: " + err.Error())
			all = append(all, "Error: 11")
		} else {
			log.Info("Getting k8s version date on "+ s.(string) +" => " + k8sVersion.GitVersion)
			all = append(all, k8sVersion.GitVersion)
		}
	}
	return all, nil
}

func GetK8sCertStatus(rows []interface{}) ([]interface{}, error) {
	var all []interface{}
	for _, s := range rows {
		cnnState, err := helper.TlsDial(s.(string) + ":6443")
		if err != nil {
			log.Error("TLS bağlantısı kurulurken bir hata oluştu: " + err.Error())
			log.Error("Sertifika verisi alınamadı =>"+ s.(string))
			all = append(all, "Error: 11")
		}else {
			expiry := cnnState().PeerCertificates[0].NotAfter
			log.Info("Getting k8s certificate expire date on "+ s.(string) + " => " + expiry.Format("02.01.2006"))
			all = append(all, expiry.Format("02.01.2006"))
		}
	}
	return all, nil
}



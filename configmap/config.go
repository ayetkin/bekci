package configmap

import (
	"bekci/model"
	log "github.com/sirupsen/logrus"
	"github.com/tkanos/gonfig"
)

var Config = model.Config{}

func init()  {
	err := gonfig.GetConf("configmap/config.json", &Config)
	if err != nil {
		log.Fatal(err)
	}
}

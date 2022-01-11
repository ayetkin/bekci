package helper

import (
	"crypto/tls"
	"encoding/json"
	"net/http"
	"time"
)

var Client = &http.Client{
	Timeout: 10 * time.Second,
	Transport: &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}},
}

func HttpGet(url string, target interface{}) error {
	r, err := Client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()
	return json.NewDecoder(r.Body).Decode(target)
}
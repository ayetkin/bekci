package model

type Clusters struct {
	Hostname string
	IP       string
	Version  string
	Date     string
	Env      string

}

type K8sVersion struct {
	GitVersion   string    `json:"gitVersion"`
}

type Config struct {
	SHEET_ID string
	CHANNEL string
	WEBHOOK_URL string
}
package sheetsapi

import (
	authentication "bekci/authentication"
	"bekci/configmap"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
)

func GetRows(rows string) ([]interface{}, error) {
	log.Warning("Getting rows from remote google sheet via api...")
	b, err := ioutil.ReadFile("authentication/credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}
	config, err := google.ConfigFromJSON(b, "https://www.googleapis.com/authentication/spreadsheets")
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}

	client := authentication.GetClient(config)
	srv, err := sheets.New(client)
	if err != nil {
		log.Fatalf("Unable to retrieve Sheets client: %v", err)
	}
	spreadsheetId := configmap.Config.SHEET_ID
	resp, err := srv.Spreadsheets.Values.Get(spreadsheetId, rows).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet: %v", err)
	}
	var all []interface{}
	if len(resp.Values) == 0 {
		log.Error("No data found.")
	} else {
		for _, row := range resp.Values {
			all = append(all, row[0])
		}
		return all, err
	}
	return nil ,err
}
package sheetsapi

import (
	"bekci/authentication"
	"bekci/configmap"
	log "github.com/sirupsen/logrus"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/sheets/v4"
	"io/ioutil"
)

func WriteRows(writeRange string, value []interface{}) {
	log.Warning("Writing values to remote Google sheet via api...")
	var values [][]interface{}
	for _, s := range value {
		var val []interface{}
		val = append(val, s)
		values = append(values, val)
	}
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

	rb := &sheets.BatchUpdateValuesRequest{
		ValueInputOption: "USER_ENTERED",
	}
	rb.Data = append(rb.Data, &sheets.ValueRange{
		Range:  writeRange,
		Values: values,
	})
	_, err = srv.Spreadsheets.Values.BatchUpdate(spreadsheetId, rb).Do()
	if err != nil {
		log.Fatal(err)
	}
}
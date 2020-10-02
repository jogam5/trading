package client

import (
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"gopkg.in/Iwark/spreadsheet.v2"
	"io/ioutil"
	"log"
)

func ConnectionGoogle(spreadsheetName string) spreadsheet.Spreadsheet {
	data, err := ioutil.ReadFile("go-cry-ade01b0b9a7c.json")
	checkError(err)
	conf, err := google.JWTConfigFromJSON(data, spreadsheet.Scope)
	checkError(err)
	client := conf.Client(context.TODO())

	service := spreadsheet.NewServiceWithClient(client)
	spreadsheet, err := service.FetchSpreadsheet(spreadsheetName)
	checkError(err)
	log.Println("####### Google ON")
	return spreadsheet
}

func checkError(err error) {
	if err != nil {
		panic(err.Error())
	}
}

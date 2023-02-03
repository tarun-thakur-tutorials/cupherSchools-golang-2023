package database

import (
	"encoding/xml"
	"log"

	"example.com/models"
	"example.com/utils"
)

const url = "https://www.ecb.europa.eu/stats/eurofxref/eurofxref-hist-90d.xml"

func LoadXMLInDatabase() {
	data, err := utils.DownloadXML(url)
	if err != nil {
		log.Fatal(err)
	}

	var envelope models.Envelope

	err = xml.Unmarshal(data, &envelope)

	if err != nil {
		log.Fatal(err)
	}

	records := utils.ConvertXmlIntoRecord(envelope)

	for _, record := range records {
		_, err := Db.Exec("INSERT INTO exchange_rates (time,currency, rate) VALUES (?,?,?)",
			record.Time,
			record.Currency,
			record.Rate,
		)
		if err != nil {
			log.Fatal(err)
		}
	}
}

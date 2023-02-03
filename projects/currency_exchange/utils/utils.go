package utils

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"example.com/models"
	"github.com/spf13/cast"
)

func DownloadXML(link string) ([]byte, error) {
	resp, err := http.Get(link)
	if err != nil {
		return []byte{}, fmt.Errorf("error during fetching the data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("Status Code: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("Error while reading the response body: %v", resp.StatusCode)
	}

	fmt.Println("the data recieved is ", string(data))

	return data, nil
}

func ConvertXmlIntoRecord(envelope models.Envelope) (records models.RecordList) {
	for _, dateDetails := range envelope.CubeList.CubeTime {
		date, err := time.Parse("2006-01-02", dateDetails.Time)
		if err != nil {
			fmt.Println(err)
			return
		}

		for _, currencyDetails := range dateDetails.CubeCurrency {
			currency := currencyDetails.Currency
			rate := cast.ToFloat64(currencyDetails.Rate)
			records = append(records, models.Record{
				Time:     date,
				Currency: currency,
				Rate:     rate,
			})
		}

	}
	fmt.Println(" the records are ", records)
	return
}

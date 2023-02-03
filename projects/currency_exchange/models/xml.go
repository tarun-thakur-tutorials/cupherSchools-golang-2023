package models

import "encoding/xml"

type Envelope struct {
	XMLName  xml.Name `xml:"Envelope"`
	Text     string   `xml:",chardata"`
	Gesmes   string   `xml:"gesmes,attr"`
	Xmlns    string   `xml:"xmlns,attr"`
	Subject  string   `xml:"subect"`
	Sender   Sender   `xml:"Sender"`
	CubeList CubeList `xml:"Cube"`
}

type Sender struct {
	Text string `xml:"chardata"`
	Name string `xml:"name"`
}

type CubeList struct {
	Text     string     `xml:",chardata"`
	CubeTime []CubeTime `xml:"Cube"`
}

type CubeTime struct {
	Text         string         `xml:",chardata"`
	Time         string         `xml:"time,attr"`
	CubeCurrency []CubeCurrency `xml:"Cube"`
}

type CubeCurrency struct {
	Text     string `xml:",chardata"`
	Currency string `xml:"currency,attr"`
	Rate     string `xml:"rate,attr"`
}

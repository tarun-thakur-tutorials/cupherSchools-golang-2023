package models

import (
	"time"
)

type RecordList = []Record

type Record struct {
	Time     time.Time `json:"time"`
	Currency string    `json:"currency"`
	Rate     float64   `json:"rate"`
}

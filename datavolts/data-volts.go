package datavolts

import (
	"time"
)

const DataVoltsTableName string = "data-volts"

type DataVolts struct {
	RTensions      []float64 `json:"rTensions"`
	STensions      []float64 `json:"sTensions"`
	TTensions      []float64 `json:"tTensions"`
	RCurrents      []float64 `json:"rCurrents"`
	SCurrents      []float64 `json:"sCurrents"`
	TCurrents      []float64 `json:"tCurrents"`
	RealTimestamp  time.Time `json:"realTimestamp"`
	QueueTimestamp time.Time `json:"queueTimestamp"`
	Timestamp      time.Time `json:"timestamp"`
	MessageID      string    `json:"messageId"`
}

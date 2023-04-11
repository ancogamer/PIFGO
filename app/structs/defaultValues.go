package structs

import (
	"strconv"
	"time"

	"github.com/ancogamer/app/structs/operation"
)

const (
	DefaultQTDYear     int    = 1
	QTDMonthYear       int    = 12
	TimeFormatYYYYMMDD string = "2006/01/02"
)

var (
	StartData  time.Time = time.Now()
	StartYear  string    = strconv.Itoa(StartData.Year())
	StartMonth string    = StartData.Month().String()
	StartDay   string    = strconv.Itoa(StartData.Day())
)

type Data struct {
	BeginValue float64 `json:"begin_value"`
	TimeWindow []operation.Year
}

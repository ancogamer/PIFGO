package structs

import (
	"strconv"
	"time"
)

const (
	DefaultQTDYear int = 1
	QTDMonthYear   int = 12
)

var (
	StartData  time.Time = time.Now()
	StartYear  string    = strconv.Itoa(StartData.Year())
	StartMonth string    = StartData.Month().String()
	StartDay   string    = strconv.Itoa(StartData.Day())
)

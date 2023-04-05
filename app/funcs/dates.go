package funcs

import (
	"strconv"
	"time"

	"github.com/ancogamer/app/structs"
	"github.com/ancogamer/app/structs/operation"
)

var getCurrentData = CurrentData()

func CurrentData() time.Time {
	return time.Now()
}

// from https://stackoverflow.com/questions/73880828/list-the-number-of-days-in-current-date-month
func DaysInMonth(t time.Time) []int {
	t = time.Date(t.Year(), t.Month(), 32, 0, 0, 0, 0, time.UTC)
	daysInMonth := 32 - t.Day()
	days := make([]int, daysInMonth)
	for i := range days {
		days[i] = i + 1
	}
	return days
}

func SetupFileCurrentTime() (timeWindow []operation.Year) {

	timeWindow = make([]operation.Year, structs.DefaultQTDYear)
	tNow := getCurrentData
	for y := 0; y < structs.DefaultQTDYear; y++ {
		timeWindow[y].Year = strconv.Itoa(tNow.Year())
		cMonth := tNow.Month()

		remainsMonths := structs.QTDMonthYear - int(cMonth)
		if int(cMonth) == 12 {
			remainsMonths = 1
		}
		// calculate the rest of the years month
		timeWindow[y].Months = make([]operation.Month, remainsMonths)
		for m := 0; m < remainsMonths; m++ {
			cMonth = tNow.Month()
			days := DaysInMonth(tNow)
			timeWindow[y].Months[m].Name = cMonth.String()
			timeWindow[y].Months[m].NumberInYear = int(cMonth)
			// since the days slice start at 0, the current quantity of days remaining on the month
			// is the the size of days - current day +1 (since we need the space for the current day)
			cDay := tNow.Day()
			qtdDaysInMonth := len(days)

			timeWindow[y].Months[m].Days = make([]operation.Day, qtdDaysInMonth-cDay+1)

			// get day of start week
			stWeekDay := tNow.Weekday()
			for d := 0; d < len(timeWindow[y].Months[m].Days); d++ {
				timeWindow[y].Months[m].Days[d].Name = stWeekDay.String()
				timeWindow[y].Months[m].Days[d].NumInMonth = int8(days[cDay-1]) // current day - 1, since slices start at 0
				stWeekDay++
				if stWeekDay == 7 {
					stWeekDay = 0
				}
				cDay++
			}
			tNow = time.Date(tNow.Year()+y, tNow.Month()+1, 01, 00, 00, 00, 00, tNow.Location())
		}
		//re calc Next Year, start on day 1
		tNow = time.Date(tNow.Year()+y, 01, 01, 00, 00, 00, 00, tNow.Location())
	}
	return
}

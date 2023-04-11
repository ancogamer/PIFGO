package funcs

import (
	"github.com/ancogamer/app/structs/operation"
)

// GetDataItem gets the position from the year month and day in the data slice, or nil case some stuff dont work
func GetPositionOperationYearItem(in *[]operation.Year, year string, month, day int) (out *[3]int) {
	if in != nil {
		for y := 0; y < len(*in); y++ {
			if (*in)[y].Year != year { // skip until next year is found
				continue
			}

			for m := 0; m < len((*in)[y].Months); m++ {
				if (*in)[y].Months[m].NumberInYear != month {
					continue
				}
				for d := 0; d < len((*in)[y].Months[m].Days); d++ {
					if int((*in)[y].Months[m].Days[d].NumInMonth) != day {
						continue
					}
					return &[3]int{y, m, d}
				}
			}

		}
	}
	return nil
}

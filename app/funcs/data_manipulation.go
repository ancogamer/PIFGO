package funcs

import (
	"github.com/ancogamer/app/structs/operation"
)

// GetDataItem gets the position from the year month and day in the data slice, or nil case some stuff dont work
func GetPositionOperationYearItem(in *[]operation.Year, year string, month, day int) (out []int) {
	if in != nil {
		for y := 0; y < len(*in); y++ { // O(n), but since the data is created ordered in ASC, still space to improve this for some really large data slices
			if (*in)[y].Year != year { // skip until next year is found
				continue
			}
			if len((*in)[y].Months) < month { // slice of months dont have the request month, case data build with current data
				return
			}
			// slice start on position 0
			if len((*in)[y].Months[month-1].Days) < day { // same reason as month
				return
			}
			return []int{y, month - 1, day - 1}
		}
	}
	return nil
}

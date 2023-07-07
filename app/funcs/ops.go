package funcs

import (
	"github.com/ancogamer/app/structs/operation"
)

// AddOrUpdateIncome adds or update in the income, IT WILL TAKE THE DAY AS THE BASE for transactions, in other words, transactions only happen in days.
func AddOrUpdateIncome(year *operation.Year, month *operation.Month, day *operation.Day, inp operation.Income) {
	// this should call updateIncomeOnMonth to update the balance on the month
}

// since i dont see any outside package usage for this im unexporting those
// updateIncomeOnMonth updates the income of a month
func updateIncomeOnMonth(in *operation.Month) {
	// this should call updateIncomeOnYear to update the balance on the year
}

// updateIncomeOnYear updates the income of the year
func updateIncomeOnYear(in *operation.Year) {

}

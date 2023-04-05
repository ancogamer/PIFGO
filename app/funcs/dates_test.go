package funcs

import (
	"testing"
	"time"

	"github.com/ancogamer/app/structs/operation"
	"github.com/google/go-cmp/cmp"
)

// go test -v -failfast -run ^TestSetupFileCurrentTime$
func TestSetupFileCurrentTime(t *testing.T) {
	tests := []struct {
		name           string
		wantTimeWindow []operation.Year
		getCurrentData time.Time
	}{
		{
			name: "ok",
			wantTimeWindow: []operation.Year{
				{
					Year: "2023",
					Months: []operation.Month{
						{Name: "December", NumberInYear: 12, Days: []operation.Day{
							{Name: "Saturday", NumInMonth: 2},
							{Name: "Sunday", NumInMonth: 3},
							{Name: "Monday", NumInMonth: 4},
							{Name: "Tuesday", NumInMonth: 5},
							{Name: "Wednesday", NumInMonth: 6},
							{Name: "Thursday", NumInMonth: 7},
							{Name: "Friday", NumInMonth: 8},
							{Name: "Saturday", NumInMonth: 9},
							{Name: "Sunday", NumInMonth: 10},
							{Name: "Monday", NumInMonth: 11},
							{Name: "Tuesday", NumInMonth: 12},
							{Name: "Wednesday", NumInMonth: 13},
							{Name: "Thursday", NumInMonth: 14},
							{Name: "Friday", NumInMonth: 15},
							{Name: "Saturday", NumInMonth: 16},
							{Name: "Sunday", NumInMonth: 17},
							{Name: "Monday", NumInMonth: 18},
							{Name: "Tuesday", NumInMonth: 19},
							{Name: "Wednesday", NumInMonth: 20},
							{Name: "Thursday", NumInMonth: 21},
							{Name: "Friday", NumInMonth: 22},
							{Name: "Saturday", NumInMonth: 23},
							{Name: "Sunday", NumInMonth: 24},
							{Name: "Monday", NumInMonth: 25},
							{Name: "Tuesday", NumInMonth: 26},
							{Name: "Wednesday", NumInMonth: 27},
							{Name: "Thursday", NumInMonth: 28},
							{Name: "Friday", NumInMonth: 29},
							{Name: "Saturday", NumInMonth: 30},
							{Name: "Sunday", NumInMonth: 31},
						}}},
				},
			},
			getCurrentData: time.Date(2023, 12, 02, 19, 03, 00, 0000, time.UTC),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getCurrentData = tt.getCurrentData
			gotTimeWindow := SetupFileCurrentTime()
			if diff := cmp.Diff(gotTimeWindow, tt.wantTimeWindow); diff != "" {
				t.Errorf("SetupFileCurrentTime() = %s", diff)
			}
		})
	}
}

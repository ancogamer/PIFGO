package funcs

import (
	"reflect"
	"testing"

	"github.com/ancogamer/app/structs/operation"
)

// go test -v -failfast -run ^TestGetPositionOperationYearItem$
func TestGetPositionOperationYearItem(t *testing.T) {
	type args struct {
		in    *[]operation.Year
		year  string
		month int
		day   int
	}
	tests := []struct {
		name    string
		args    args
		wantOut *[3]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := GetPositionOperationYearItem(tt.args.in, tt.args.year, tt.args.month, tt.args.day); !reflect.DeepEqual(gotOut, tt.wantOut) {
				t.Errorf("GetPositionOperationYearItem() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

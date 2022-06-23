package data

import (
	"reflect"
	"testing"
)

var testData = map[string]dailyData{
	"2022-01-03": {"2", "3", "1", "4", "10"},
	"2022-01-04": {"3", "4", "2", "5", "11"},
	"2022-01-05": {"4", "5", "3", "6", "12"},
	"2022-01-06": {"5", "6", "4", "7", "13"},
	"2022-01-07": {"6", "7", "5", "8", "14"},
	"2022-01-08": {"7", "8", "6", "9", "15"},
	"2022-06-23": {"9999", "9999", "9999", "9999", "9999"},
}

func Test_getDates(t *testing.T) {
	type args struct {
		timeseries map[string]dailyData
		nDays      int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"give me a day", args{testData, 1}, []string{"2022-06-23"}},
		{"give me two days", args{testData, 2}, []string{"2022-06-23", "2022-01-08"}},
		{"give me three days", args{testData, 3}, []string{"2022-06-23", "2022-01-08", "2022-01-07"}},
		{"i just need a week", args{testData, 7}, []string{"2022-06-23", "2022-01-08", "2022-01-07", "2022-01-06", "2022-01-05", "2022-01-04", "2022-01-03"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getDates(tt.args.timeseries, tt.args.nDays); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getDates() = %v, want %v", got, tt.want)
			}
		})
	}
}

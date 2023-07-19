package core

import (
	"reflect"
	"testing"
)

func Test_chargeJsonData(t *testing.T) {
	type args struct {
		dataPath string
	}
	tests := []struct {
		name string
		args args
		want Elements
	}{
		{
			name: "one",
			args: args{"data.json"},
			want: []map[string]interface{}{
				{
					"Name": "ro",
					"Tm":   "rogerDoc",
				},
				{
					"Name": "evelyn <br> <br> something different ",
					"Tm":   "eveDoc",
				},
			},
		},
		{
			name: "one",
			args: args{"data.csv"},
			want: []map[string]interface{}{
				{
					"Name": "ro",
					"Tm":   "rogerDoc",
				},
				{
					"Name": "evelyn <br> <br> something different ",
					"Tm":   "eveDoc",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := chargeData(tt.args.dataPath)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargeData() = %v, want %v", got, tt.want)
			}
		})
	}
}

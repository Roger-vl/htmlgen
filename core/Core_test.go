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
		want map[string]interface{}
	}{
		{name: "one", args: args{"data.json"}, want: map[string]interface{}{"name": "roger"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := chargeJsonData(tt.args.dataPath); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("chargeJsonData() = %v, want %v", got, tt.want)
			}
		})
	}
}

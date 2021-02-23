package card

import (
	"reflect"
	"testing"
)

func TestMapRowToTransaction(t *testing.T) {
	transactions := []Transaction{
		{
			UserId: 0,
			Sum:    1000,
			MCC:    "5411",
		},
		{
			UserId: 1,
			Sum:    2000,
			MCC:    "5812",
		},
	}
	rows := [][]string{{"0", "1000", "5411"}, {"1", "2000", "5812"}}
	type args struct {
		m [][]string
	}
	tests := []struct {
		name string
		args args
		want []Transaction
	}{
		{
			name: "map to struct",
			args: args {
				rows,
			},
			want: transactions,
		},
	}
		for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MapRowToTransaction(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MapRowToTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
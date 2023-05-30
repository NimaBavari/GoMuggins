package muggins

import (
	"reflect"
	"testing"
)

func TestGetMin(t *testing.T) {
	type args struct {
		x T
		y T
	}
	tests := []struct {
		name string
		args args
		want T
	}{
		{
			name: "integer positive same",
			args: args{2, 2},
			want: 2,
		}, {
			name: "integer positive not same",
			args: args{3, 4},
			want: 3,
		}, {
			name: "integer mixed",
			args: args{5, -2},
			want: -2,
		}, {
			name: "integer negative not same",
			args: args{-4, -8},
			want: -8,
		}, {
			name: "integer negative same",
			args: args{-6, -6},
			want: -6,
		}, {
			name: "float positive same",
			args: args{5.2, 5.2},
			want: 5.2,
		}, {
			name: "float positive not same",
			args: args{4.3, 7.4},
			want: 4.3,
		}, {
			name: "float mixed",
			args: args{-2.6, 3.6},
			want: -2.6,
		}, {
			name: "float negative not same",
			args: args{-1.7, -2.7},
			want: -2.7,
		}, {
			name: "float negative same",
			args: args{-3.2, -3.2},
			want: -3.2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMin(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetMin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHas(t *testing.T) {
	type args struct {
		arr  []T
		elem T
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "containing",
			args: args{[]int{2, 3, 7}, 7},
			want: true,
		}, {
			name: "not containing",
			args: args{[]int{4, 16, 32}, 3},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Has(tt.args.arr, tt.args.elem); got != tt.want {
				t.Errorf("Has() = %v, want %v", got, tt.want)
			}
		})
	}
}

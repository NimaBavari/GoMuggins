package muggins

import (
	"reflect"
	"testing"
)

func TestTile_Face(t *testing.T) {
	type fields struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name:   "blank",
			fields: fields{0, 0},
			want:   0,
		}, {
			name:   "six-two",
			fields: fields{6, 2},
			want:   8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tile{
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := tr.Face(); got != tt.want {
				t.Errorf("Tile.Face() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_IsDouble(t *testing.T) {
	type fields struct {
		left  int
		right int
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{
			name:   "fours",
			fields: fields{4, 4},
			want:   true,
		}, {
			name:   "two-five",
			fields: fields{2, 5},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tile{
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := tr.IsDouble(); got != tt.want {
				t.Errorf("Tile.IsDouble() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_IsSame(t *testing.T) {
	type fields struct {
		left  int
		right int
	}
	type args struct {
		o Tile
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "same",
			fields: fields{3, 4},
			args:   args{Tile{4, 3}},
			want:   true,
		}, {
			name:   "not same",
			fields: fields{2, 0},
			args:   args{Tile{1, 4}},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tile{
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := tr.IsSame(tt.args.o); got != tt.want {
				t.Errorf("Tile.IsSame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTile_IsPlayable(t *testing.T) {
	type fields struct {
		left  int
		right int
	}
	type args struct {
		e int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			name:   "playable",
			fields: fields{2, 4},
			args:   args{4},
			want:   true,
		}, {
			name:   "non-playable",
			fields: fields{3, 4},
			args:   args{2},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := Tile{
				left:  tt.fields.left,
				right: tt.fields.right,
			}
			if got := tr.IsPlayable(tt.args.e); got != tt.want {
				t.Errorf("Tile.IsPlayable() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFromString(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want Tile
	}{
		{
			name: "three-five",
			args: args{"3:5"},
			want: Tile{3, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FromString(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromString() = %v, want %v", got, tt.want)
			}
		})
	}
}

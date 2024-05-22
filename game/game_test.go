package game

import (
	"reflect"
	"testing"
)

func TestIsWhite(t *testing.T) {
	type args struct {
		b       GameGo
		counter int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "Test White Turn",
			args: args{
				b:       &gameGo{}, // Simulate a gameGo
				counter: 0,         // Simulate the first move (even counter)
			},
			want: true,
		},
		{
			name: "Test Black Turn",
			args: args{
				b:       &gameGo{}, // Simulate a gameGo
				counter: 1,         // Simulate the second move (odd counter)
			},
			want: false,
		}, {
			name: "Test Black Turn after 35 moves",
			args: args{
				b:       &gameGo{}, // Simulate a gameGo
				counter: 35,        // Simulate the second move (odd counter)
			},
			want: false,
		}, {
			name: "Test White Turn after 100 moves",
			args: args{
				b:       &gameGo{}, // Simulate a gameGo
				counter: 100,       // Simulate the second move (odd counter)
			},
			want: true,
		},
		// Add more test cases as needed
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsWhite(tt.args.b, tt.args.counter); got != tt.want {
				t.Errorf("IsWhite() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNewBoard(t *testing.T) {
	tests := []struct {
		name string
		want GameGo
	}{
		{name: "Test new gameGo",
			want: &gameGo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewGame(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewGame() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gameGo_CheckTurn(t *testing.T) {
	type fields struct {
		turn Turn
	}
	type args struct {
		counter int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Turn
	}{{
		name:   "test_CheckTurn",
		fields: fields{turn: Black},
		args:   args{counter: 10},
		want:   White,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &gameGo{
				turn: tt.fields.turn,
			}
			if got := b.CheckTurn(tt.args.counter); got != tt.want {
				t.Errorf("CheckTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

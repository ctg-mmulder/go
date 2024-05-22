package models

import (
	"reflect"
	"testing"
)

func TestIsWhite(t *testing.T) {
	type args struct {
		b       Board
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
				b:       &board{}, // Simulate a board
				counter: 0,        // Simulate the first move (even counter)
			},
			want: true,
		},
		{
			name: "Test Black Turn",
			args: args{
				b:       &board{}, // Simulate a board
				counter: 1,        // Simulate the second move (odd counter)
			},
			want: false,
		}, {
			name: "Test Black Turn after 35 moves",
			args: args{
				b:       &board{}, // Simulate a board
				counter: 35,       // Simulate the second move (odd counter)
			},
			want: false,
		}, {
			name: "Test White Turn after 100 moves",
			args: args{
				b:       &board{}, // Simulate a board
				counter: 100,      // Simulate the second move (odd counter)
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
		want Board
	}{
		{name: "Test new board",
			want: &board{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewBoard(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_board_CheckTurn(t *testing.T) {
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
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &board{
				turn: tt.fields.turn,
			}
			if got := b.CheckTurn(tt.args.counter); got != tt.want {
				t.Errorf("CheckTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

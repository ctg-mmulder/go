package game

import (
	"reflect"
	"testing"
)

func TestIsWhite(t *testing.T) {
	tests := []struct {
		name  string
		game  GameGo
		plays int
		want  bool
	}{
		{
			name:  "first play Black Turn ",
			game:  NewGame(),
			plays: 1,
			want:  false,
		},
		{
			name:  "Second Play White Turn",
			game:  NewGame(),
			plays: 2,
			want:  true,
		},
		{
			name:  "third play Black Turn ",
			game:  NewGame(),
			plays: 101,
			want:  false,
		},
		{
			name:  "Fourth Play White Turn",
			game:  NewGame(),
			plays: 36,
			want:  true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			for i := 0; i < tt.plays; i++ {
				tt.game.UpdateTurnCounter()
			}
			if got := tt.game.IsWhiteTurn(); got != tt.want {
				t.Errorf("IsWhiteTurn() = %v, want %v", got, tt.want)
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
			for i := 0; i < tt.args.counter; i++ {
				b.UpdateTurnCounter()
			}
			if got := b.CheckTurn(); got != tt.want {
				t.Errorf("CheckTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

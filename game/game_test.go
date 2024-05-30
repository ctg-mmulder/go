package game

import (
	"github.com/go-go/game/models"
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
				tt.game.PlayTile(50, 50)
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
			want: &gameGo{turn: White, turnCounter: 0, intersects: createIntersects()},
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
		turn           Turn
		intersectstest []models.Intersect
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
		fields: fields{turn: Black, intersectstest: createTestIntersects()},
		args:   args{counter: 10},
		want:   White,
	},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := &gameGo{
				turn: tt.fields.turn, intersects: tt.fields.intersectstest,
			}
			for i := 0; i < tt.args.counter; i++ {
				b.PlayTile(0, 0)
			}
			if got := b.CheckTurn(); got != tt.want {
				t.Errorf("CheckTurn() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_gameGo_PlayTile(t *testing.T) {
	type fields struct {
		game GameGo
	}
	type args struct {
		x           float64
		y           float64
		turncounter int
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   models.State
	}{{
		name: "test_PlayTile_unevenTurn_black",
		fields: fields{
			game: NewGame(),
		},
		args: args{
			x:           50,
			y:           50,
			turncounter: 49},
		want: models.Black,
	}, {
		name: "test_PlayTile_evenTurn_White",
		fields: fields{
			game: NewGame(),
		},
		args: args{
			x:           50,
			y:           50,
			turncounter: 66},
		want: models.White,
	},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.fields.game.SetCounter(tt.args.turncounter)
			tt.fields.game.PlayTile(tt.args.x, tt.args.y)

			if got := tt.fields.game.IntersectState(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("IntersectState() = %v, want %v", got, tt.want)
			}
		})
	}
}
func createTestIntersects() []models.Intersect {
	var intersects []models.Intersect
	intersect := models.Intersect{
		Location: models.Location{X: float64(0), Y: float64(0)},
		State:    models.Free,
	}
	intersects = append(intersects, intersect)
	return intersects
}

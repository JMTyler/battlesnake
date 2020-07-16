package db

import (
	"fmt"
	snek "github.com/JMTyler/battlesnake/_"
	"time"
)

type Frame struct {
	tableName struct{} `pg:"\"Frames\""`

	ID        int64     `pg:", notnull"`
	CreatedAt time.Time `pg:", notnull, default:now()"`
	UpdatedAt time.Time `pg:", notnull, default:now()"`

	Important bool   `pg:", notnull, default:false"`
	GameID    string `pg:", notnull"`
	SnakeID   string `pg:", notnull"`
	Name      string `pg:", notnull"`
	Turn      int    `pg:", notnull, use_zero"`
	Move      string `pg:","`
	Duration  int64  `pg:","`

	Context snek.Context `pg:","`
}

type frameKey struct {
	GameID  string
	SnakeID string
	// TODO: Nothing seems to enable .WhereStruct() to accept Turn as a zero. 😩
	Turn int `pg:", required, notnull, use_zero"`
}

func NewFrame(context snek.Context) *Frame {
	return &Frame{
		GameID:  context.Game.ID,
		SnakeID: context.You.ID,
		Name:    context.You.Name,
		Turn:    context.Turn,
		Context: context,
	}
}

func (f *Frame) getKey() frameKey {
	return frameKey{f.GameID, f.SnakeID, f.Turn}
}

func (f *Frame) Insert() {
	NOW := time.Now()
	f.CreatedAt = NOW
	f.UpdatedAt = NOW

	if _, err := DB.Model(f).OnConflict("DO NOTHING").Insert(f); err != nil {
		panic(err)
	}
}

func (f *Frame) Update(move string, duration int64) {
	if _, err := DB.Model(f).WhereStruct(f.getKey()).Set("move = ?, duration = ?, updated_at = ?", move, duration, time.Now()).Update(); err != nil {
		panic(err)
	}
}

func (f *Frame) String() string {
	return fmt.Sprintf("Frame{game:%v, snake:%v, turn:%v -> move:%v}", f.GameID, f.SnakeID, f.Turn, f.Move)
}

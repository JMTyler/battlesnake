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

func PruneGames() {
	for {
		numRows, err := DB.Model(&Frame{}).Count()
		if err != nil {
			panic(err)
		}

		if numRows < 10000 {
			return
		}

		// Find the oldest game in the database.
		var gameID string
		if err := DB.Model(&Frame{}).Column("game_id").Limit(1).Order("created_at ASC").Where("important = ?", false).Select(&gameID); err != nil {
			panic(err)
		}

		// And delete it.
		if _, err := DB.Model(&Frame{}).Where("game_id = ?", gameID).Delete(); err != nil {
			panic(err)
		}
	}
}

func (f *Frame) String() string {
	return fmt.Sprintf("Frame{game:%q, snake:%q, turn:%v -> move:%v}", f.GameID, f.SnakeID, f.Turn, f.Move)
}

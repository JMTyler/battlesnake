package battlesnake

type Context struct {
	Turn  int    `json:"turn"`
	Game  *Game  `json:"game"`
	You   *Snake `json:"you"`
	Board *Board `json:"board"`
}

type Game struct {
	ID      string   `json:"id"`
	Timeout int      `json:"timeout"`
	Ruleset *Ruleset `json:"ruleset"`
	Dev     bool     `json:"dev"`
}

type Ruleset struct {
	Name string `json:"name"`
}

func (ctx *Context) Prepare() {
	ctx.Board.Prepare(ctx)
}

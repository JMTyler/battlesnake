package battlesnake

type Context struct {
	Turn  int    `json:"turn"`
	Game  *Game  `json:"game"`
	You   *Snake `json:"you"`
	Board *Board `json:"board"`
}

type Game struct {
	ID      string `json:"id"`
	Timeout int    `json:"timeout"`
	Dev     bool   `json:"dev"`
}

func (ctx *Context) Prepare() {
	ctx.Board.Prepare(ctx)
}

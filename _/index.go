package tactics

type TacticOptions struct {
	Health       int
	Distance     int
	Advantage    int
	Disadvantage int
}

type Position struct {
	X int
	Y int
}

type Context struct {
	Turn  int
	Game  Game
	You   Snake
	Board Board
}

type State struct {
	Move   string
	Snakes map[string]SnakeState
}

type Game struct {
	ID      string
	Timeout int
	Dev     bool
}

type Snake struct {
	ID     string
	Head   Position
	Body   []Position
	Length int
	Health int
}

type Board struct {
	Width  int
	Height int
	Snakes []Snake
	Food   []Position
}

type SnakeState struct {
	Move string
}

//const _    = require('lodash');
//const fs   = require('fs');
//const path = require('path');
//
//const thisFile = path.basename(__filename);
//const files = fs.readdirSync(__dirname).filter((file) => (file !== thisFile));
//const tactics = _.map(files, (file) => require(`./${file}`));
//
//module.exports = Object.assign({}, ...tactics);

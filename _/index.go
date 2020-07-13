package tactics

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
	Move string
}

type Game struct {
	GameID  string
	Timeout int
	Dev     bool
}

type Snake struct {
	Head   Position
	Body   []Position
	Length int
}

type Board struct {
	Width  int
	Height int
	Snakes []Snake
	Food   []Position
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

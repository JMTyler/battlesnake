const _ = require('lodash');

const GetAdjacentTiles = ({ x, y }) => {
	return {
		up:    { x, y: y + 1 },
		down:  { x, y: y - 1 },
		left:  { x: x - 1, y },
		right: { x: x + 1, y },
	};
};

const IsOutsideBoard = ({ x, y }, board) => {
	return x < 0 || y < 0 || x >= board.width || y >= board.height;
};

const IsDeadly = (pos, { board }) => {
	if (IsOutsideBoard(pos, board)) {
		return true;
	}

	const anySnakeCollision = _.some(board.snakes, (snake) => {
		// TODO: Only drop the tail piece if the snake HASN'T just eaten a disc.
		// TODO: Or, drop the tail either way, but consider that spot risky.
		return _.some(_.initial(snake.body), pos);
	});
	if (anySnakeCollision) {
		return true;
	}

	return false;
};

const IsRisky = (pos, { board, you }) => {
	return _.some(board.snakes, (snake) => {
		if (snake.id === you.id) return false;
		const adjacent = GetAdjacentTiles(snake.head);
		return _.some(adjacent, pos);
	});
};

const IsSafe = (pos, context) => {
	return !IsDeadly(pos, context) && !IsRisky(pos, context);
};

module.exports = {
	GetAdjacentTiles,
	IsOutsideBoard,
	IsDeadly,
	IsRisky,
	IsSafe,
};

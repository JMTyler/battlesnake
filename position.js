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

const IsDeadly = (pos, { board, you }) => {
	if (IsOutsideBoard(pos, board)) {
		return true;
	}

	const anySnakeCollision = _.some(board.snakes, (snake) => {
		// TODO: Only drop the tail piece if the snake HASN'T just eaten a disc.
		// TODO: Or, drop the tail either way, but consider that spot risky.
		const collision = _.some(_.initial(snake.body), pos);
		if (!collision) {
			return false;
		}

		if (Matches(pos, snake.head) && you.length > snake.length) {
			return false;
		}

		return true;
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
		const gettinSpicy = _.some(adjacent, pos);

		return gettinSpicy && you.length <= snake.length;
	});
};

const IsSafe = (pos, context) => {
	return !IsDeadly(pos, context) && !IsRisky(pos, context);
};

const Matches = (posA, posB) => {
	return _.isEqual(posA, posB);
};

const ToDirection = (to, from) => {
	const x = Math.sign(to.x - from.x);
	const y = Math.sign(to.y - from.y);

	if (x !== 0 && y !== 0) {
		return null;
	}

	if (x !== 0) {
		return x > 0 ? 'right' : 'left';
	}

	return y > 0 ? 'up' : 'down';
};

module.exports = {
	GetAdjacentTiles,
	IsOutsideBoard,
	IsDeadly,
	IsRisky,
	IsSafe,
	Matches,
	ToDirection,
};

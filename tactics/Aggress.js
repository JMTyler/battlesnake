const _ = require('lodash');

const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const Aggress = ({ context, adjacent }) => {
	// TODO: Find not just the first prey, but the closest one.
	const prey = _.find(context.board.snakes, (snake) => {
		return context.you.length >= snake.length + 2;
	});

	if (!prey) {
		return false;
	}

	const move = pathfinding.ApproachTarget(prey.head, context);
	utils.LogMove(context.turn, move, 'fired up, ready to kill');

	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { Aggress };

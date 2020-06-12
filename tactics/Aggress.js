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
	
	const targetOptions = _.filter(position.GetAdjacentTiles(prey.head), (pos) => position.IsSafe(pos, context));
	const target = pathfinding.FindClosestTarget(context.you.head, targetOptions);
	const move = pathfinding.ApproachTarget(target, context);
	utils.LogMove(context.turn, move, 'Aggress');

	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { Aggress };

const _ = require('lodash');

const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const Aggress = ({ context, adjacent }) => {
	const preyOptions = _.filter(context.board.snakes, (snake) => (context.you.length >= snake.length + 2));
	if (_.isEmpty(preyOptions)) {
		return false;
	}

	const prey = (preyOptions.length === 1) ? preyOptions[0] : pathfinding.FindClosestTarget(context.you.head, preyOptions);
	const targetOptions = _.filter(position.GetAdjacentTiles(prey.head), (pos) => position.IsSafe(pos, context));
	if (_.isEmpty(targetOptions)) {
		return false;
	}

	const target = pathfinding.FindClosestTarget(context.you.head, targetOptions);
	const move = pathfinding.ApproachTarget(target, context);
	utils.LogMove(context.turn, move, 'Aggress');

	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { Aggress };

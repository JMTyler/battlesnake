const _ = require('lodash');

const movement = require('../movement');
const position = require('../position');
const utils    = require('../utils');

const Aggrieve = ({ context, adjacent }) => {
	const preyOptions = _.filter(context.board.snakes, (snake) => (context.you.length >= snake.length + 2));
	if (_.isEmpty(preyOptions)) {
		return false;
	}

	const prey = movement.FindClosestTarget(context.you.head, _.map(preyOptions, 'head'));
	const targetOptions = _.filter(position.GetAdjacentTiles(prey), (pos) => position.IsSafe(pos, context));
	if (_.isEmpty(targetOptions)) {
		return false;
	}

	const target = movement.FindClosestTarget(context.you.head, targetOptions);
	const move = movement.ApproachTarget(target, context);
	utils.LogMove(context.turn, move, 'Aggrieve');

	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { Aggrieve };

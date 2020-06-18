const _ = require('lodash');

const movement = require('../movement');
const position = require('../position');
const utils    = require('../utils');

const Abscond = ({ context, adjacent }) => {
	const predators = _.filter(context.board.snakes, (snake) => (snake.length > context.you.length));
	if (_.isEmpty(predators)) {
		return false;
	}

	const predator = movement.FindClosestTarget(context.you.head, _.map(predators, 'head'));
	const distance = movement.GetDistance(context.you.head, predator);
	if (distance > 3) {
		return false;
	}

	const vector = movement.GetVector(context.you.head, predator);
	const escapeVector = {
		x : -1 * vector.weight.x,
		y : -1 * vector.weight.y,
	};
	const escapeTarget = {
		x : _.clamp(escapeVector.x + context.you.head.x, 0, context.board.width - 1),
		y : _.clamp(escapeVector.y + context.you.head.y, 0, context.board.height - 1),
	};

	const move = movement.ApproachTarget(escapeTarget, context);
	utils.LogMove(context.turn, move, 'Abscond');

	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { Abscond };

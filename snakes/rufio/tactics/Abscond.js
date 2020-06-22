const _ = require('lodash');

const movement = require('../../../movement');

const Abscond = (options = {}) => {
	options = Object.assign({
		disadvantage : 1,
		distance     : Infinity,
	}, options);

	return ({ context }) => {
		const predators = _.filter(context.board.snakes, (snake) => (context.you.length <= snake.length - options.disadvantage));
		if (_.isEmpty(predators)) {
			return false;
		}

		const predator = movement.FindClosestTarget(context.you.head, _.map(predators, 'head'));
		const distanceToPredator = movement.GetDistance(context.you.head, predator);
		if (distanceToPredator > options.distance) {
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

		return movement.ApproachTarget(escapeTarget, context);
	};
};

module.exports = { Abscond };

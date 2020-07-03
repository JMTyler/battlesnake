const board    = require('../../../board');
const movement = require('../../../movement');

const Eat = (options = {}) => {
	options = Object.assign({
		health   : Infinity,
		distance : Infinity,
	}, options);

	return ({ context }) => {
		if (context.you.health > options.health) {
			return false;
		}

		const food = board.FindClosestFood(context);
		if (!food) {
			return false;
		}

		const distanceToFood = movement.GetDistance(context.you.head, food);
		if (distanceToFood > options.distance) {
			return false;
		}

		return movement.ApproachTarget(food, context);
	};
};

module.exports = { Eat };

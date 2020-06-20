const board    = require('../../../board');
const movement = require('../../../movement');

const Eat = ({ health = Infinity, distance = Infinity }) => {
	return ({ context }) => {
		if (context.you.health > health) {
			return false;
		}

		const food = board.FindClosestFood(context);
		const distanceToFood = movement.GetDistance(context.you.head, food);
		if (distanceToFood > distance) {
			return false;
		}

		return movement.ApproachTarget(food, context);
	};
};

module.exports = { Eat };

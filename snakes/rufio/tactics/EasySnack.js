const board    = require('../../../board');
const movement = require('../../../movement');

const EasySnack = ({ distance = Infinity }) => {
	return ({ context }) => {
		const closestFood = board.FindClosestFood(context);
		const distanceToFood = movement.GetDistance(context.you.head, closestFood);
		if (distanceToFood > distance) {
			return false;
		}

		return movement.ApproachTarget(closestFood, context);
	};
};

module.exports = { EasySnack };

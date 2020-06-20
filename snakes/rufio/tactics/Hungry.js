const board    = require('../../../board');
const movement = require('../../../movement');

const Hungry = ({ health = Infinity }) => {
	return ({ context }) => {
		if (context.you.health > health) {
			return false;
		}

		return movement.ApproachTarget(board.FindClosestFood(context), context);
	};
};

module.exports = { Hungry };

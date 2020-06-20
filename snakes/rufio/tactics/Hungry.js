const board    = require('../../../board');
const movement = require('../../../movement');
const utils    = require('../../../utils');

const Hungry = ({ health = Infinity }) => {
	return ({ context }) => {
		if (context.you.health > health) {
			return false;
		}

		const move = movement.ApproachTarget(board.FindClosestFood(context), context);
		utils.LogMove(context.turn, move, 'Hungry');
		return move;
	};
};

module.exports = { Hungry };

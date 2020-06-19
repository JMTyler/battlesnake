const board    = require('../board');
const movement = require('../movement');
const position = require('../position');
const utils    = require('../utils');

const Hungry = ({ health = Infinity }) => {
	return ({ context, adjacent }) => {
		if (context.you.health > health) {
			return false;
		}

		const move = movement.ApproachTarget(board.FindClosestFood(context), context);
		utils.LogMove(context.turn, move, 'Hungry');
		const isSafe = position.IsSafe(adjacent[move], context);
		return isSafe && move;
	};
};

module.exports = { Hungry };

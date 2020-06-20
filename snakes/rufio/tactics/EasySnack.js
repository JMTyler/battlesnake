const board    = require('../../../board');
const movement = require('../../../movement');
const position = require('../../../position');
const utils    = require('../../../utils');

const EasySnack = ({ distance = Infinity }) => {
	return ({ context, adjacent }) => {
		const closestFood = board.FindClosestFood(context);
		const distanceToFood = movement.GetDistance(context.you.head, closestFood);
		if (distanceToFood > distance) {
			return false;
		}

		const move = movement.ApproachTarget(closestFood, context);
		utils.LogMove(context.turn, move, 'Easy Snack');
		const isSafe = position.IsSafe(adjacent[move], context);
		return isSafe && move;
	};
};

module.exports = { EasySnack };

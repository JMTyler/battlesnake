const board       = require('../board');
const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const EasySnack = ({ context, adjacent }) => {
	const closestFood = board.FindClosestFood(context);
	const distanceToFood = pathfinding.GetDistance(context.you.head, closestFood);
	if (distanceToFood > 2) {
		return false;
	}
	
	const move = pathfinding.ApproachTarget(closestFood, context);
	utils.LogMove(context.turn, move, 'Easy Snack');
	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { EasySnack };

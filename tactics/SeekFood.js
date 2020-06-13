const board       = require('../board');
const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const SeekFood = ({ context, adjacent }) => {
	const closestFood = board.FindClosestFood(context);
	const distanceToFood = pathfinding.GetDistance(context.you.head, closestFood);
	if (distanceToFood > 2 && context.you.health > 95) {
		return false;
	}
	
	const move = pathfinding.ApproachTarget(closestFood, context);
	utils.LogMove(context.turn, move, 'Seek Food');
	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { SeekFood };

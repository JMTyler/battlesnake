const board       = require('../board');
const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const SeekFood = ({ context, adjacent }) => {
	if (context.you.health > 91) {
		return false;
	}
	
	const move = pathfinding.ApproachTarget(board.FindClosestFood(context), context);
	utils.LogMove(context.turn, move, 'Seek Food');
	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { SeekFood };

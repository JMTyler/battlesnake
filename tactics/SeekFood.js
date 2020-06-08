const board       = require('../board');
const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const SeekFood = ({ context, adjacent }) => {
	const maxTravel = context.board.width + context.board.height - 2;
	if (context.you.health > maxTravel) {
		return false;
	}
	
	const move = pathfinding.ApproachTarget(board.FindClosestFood(context), context);
	utils.LogMove(context.turn, move, 'hungry, seeking food');
	const isSafe = position.IsSafe(adjacent[move], context);
	return move;
};

module.exports = { SeekFood };

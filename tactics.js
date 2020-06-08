const _ = require('lodash');

const board       = require('./board');
const pathfinding = require('./pathfinding');
const position    = require('./position');
const utils       = require('./utils');

const Continue = ({ context, state, adjacent }) => {
	if (context.turn === 0) utils.LogMove(context.turn, state.move, 'first turn');
	//else utils.LogMove(context.turn, state.move, 'no change');

	const isSafe = position.IsSafe(adjacent[state.move], context);

	return isSafe && state.move;
};

const SeekFood = ({ context, state, adjacent }) => {
	if (context.you.health > state.maxTravel) {
		return false;
	}
	
	const move = pathfinding.ApproachTarget(board.FindClosestFood(context), context);
	utils.LogMove(context.turn, move, 'hungry, seeking food');
	return move;
};

const SeekTail = ({ context, adjacent }) => {
	const move = pathfinding.ApproachTarget(_.last(context.you.body), context);
	utils.LogMove(context.turn, move, 'unsafe, approaching tail');
	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

const RotateUntilSafe = ({ context, state, adjacent }) => {
	const rotate = { right: 'down', down: 'left', left: 'up', up: 'right' };

	let turns = 0;
	let isSafe = false;
	let move = state.move;
	do {
		move = rotate[move];
		turns += 1;
		isSafe = position.IsSafe(adjacent[move], context);
		utils.LogMove(context.turn, move, 'still unsafe, had to turn');
	} while (turns < 4 && !isSafe);

	return isSafe && move;
};

module.exports = {
	Continue,
	SeekFood,
	SeekTail,
	RotateUntilSafe,
};

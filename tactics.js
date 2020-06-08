const _ = require('lodash');

const board       = require('./board');
const pathfinding = require('./pathfinding');
const position    = require('./position');
const utils       = require('./utils');

const Continue = (move, context) => {
	if (context.turn === 0) utils.LogMove(context.turn, move, 'first turn');
	//else utils.LogMove(context.turn, state.move, 'no change');

	const adjacent = position.GetAdjacentTiles(context.you.head);
	const isSafe = position.IsSafe(adjacent[move], context);

	return [isSafe, move];
};

const SeekFood = (context) => {
	const adjacent = position.GetAdjacentTiles(context.you.head);
	const move = pathfinding.ApproachTarget(board.FindClosestFood(context), context);
	utils.LogMove(context.turn, move, 'hungry, seeking food');
	const isSafe = position.IsSafe(adjacent[move], context);
	return [isSafe, move];
};

const SeekTail = (context) => {
	const adjacent = position.GetAdjacentTiles(context.you.head);
	const move = pathfinding.ApproachTarget(_.last(context.you.body), context);
	utils.LogMove(context.turn, move, 'unsafe, approaching tail');
	const isSafe = position.IsSafe(adjacent[move], context);
	return [isSafe, move];
};

const RotateUntilSafe = (move, context) => {
	const adjacent = position.GetAdjacentTiles(context.you.head);
	const rotate = { right: 'down', down: 'left', left: 'up', up: 'right' };

	let turns = 0;
	let isSafe = false;
	do {
		move = rotate[move];
		turns += 1;
		isSafe = position.IsSafe(adjacent[move], context);
		utils.LogMove(context.turn, move, 'still unsafe, had to turn');
	} while (turns < 4 && !isSafe);

	return [isSafe, move];
};

module.exports = {
	Continue,
	SeekFood,
	SeekTail,
	RotateUntilSafe,
};

const position = require('../position');
const utils    = require('../utils');

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

module.exports = { RotateUntilSafe };

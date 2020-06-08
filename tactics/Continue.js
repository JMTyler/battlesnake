const position = require('../position');
const utils    = require('../utils');

const Continue = ({ context, state, adjacent }) => {
	if (context.turn === 0) utils.LogMove(context.turn, state.move, 'first turn');
	//else utils.LogMove(context.turn, state.move, 'no change');

	const isSafe = position.IsSafe(adjacent[state.move], context);

	return isSafe && state.move;
};

module.exports = { Continue };

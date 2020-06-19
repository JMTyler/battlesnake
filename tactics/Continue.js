const position = require('../position');
const utils    = require('../utils');

const Continue = () => {
	return ({ context, state, adjacent }) => {
		if (context.turn === 0) utils.LogMove(context.turn, state.move, 'Initial Move');
		else utils.LogMove(context.turn, state.move, 'Continue');

		const isSafe = position.IsSafe(adjacent[state.move], context);

		return isSafe && state.move;
	};
};

module.exports = { Continue };

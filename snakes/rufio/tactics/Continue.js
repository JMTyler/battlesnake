const utils = require('../../../utils');

const Continue = () => {
	return ({ context, state }) => {
//		if (context.turn === 0) utils.LogMove(context.turn, state.move, 'Initial Move');
//		else utils.LogMove(context.turn, state.move, 'Continue');

		return state.move;
	};
};

module.exports = { Continue };

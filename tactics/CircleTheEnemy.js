const position = require('../position');

const { Aggrieve } = require('./Aggrieve');

const CircleTheEnemy = ({ context, adjacent }) => {
	const move = Aggrieve({ context, adjacent });
	if (!move) {
		return false;
	}

	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { CircleTheEnemy };

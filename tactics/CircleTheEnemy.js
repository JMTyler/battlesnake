const position = require('../position');

const { Aggress } = require('./Aggress');

const CircleTheEnemy = ({ context, adjacent }) => {
	const move = Aggress({ context, adjacent });
	if (!move) {
		return false;
	}

	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { CircleTheEnemy };

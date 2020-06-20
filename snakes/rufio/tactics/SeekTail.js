const _ = require('lodash');

const movement = require('../../../movement');
const position = require('../../../position');
const utils    = require('../../../utils');

const SeekTail = () => {
	return ({ context, adjacent }) => {
		const move = movement.ApproachTarget(_.last(context.you.body), context);
		utils.LogMove(context.turn, move, 'Seek Tail');
		const isSafe = position.IsSafe(adjacent[move], context);
		return isSafe && move;
	};
};

module.exports = { SeekTail };

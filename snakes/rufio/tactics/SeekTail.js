const _ = require('lodash');

const movement = require('../../../movement');
const utils    = require('../../../utils');

const SeekTail = () => {
	return ({ context }) => {
		const move = movement.ApproachTarget(_.last(context.you.body), context);
		utils.LogMove(context.turn, move, 'Seek Tail');
		return move;
	};
};

module.exports = { SeekTail };

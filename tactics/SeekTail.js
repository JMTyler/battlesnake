const _ = require('lodash');

const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const SeekTail = ({ context, adjacent }) => {
	const move = pathfinding.ApproachTarget(_.last(context.you.body), context);
	utils.LogMove(context.turn, move, 'Seek Tail');
	const isSafe = position.IsSafe(adjacent[move], context);
	return isSafe && move;
};

module.exports = { SeekTail };

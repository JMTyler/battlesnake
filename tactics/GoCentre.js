const _ = require('lodash');

const pathfinding = require('../pathfinding');
const position    = require('../position');
const utils       = require('../utils');

const GoCentre = ({ context, adjacent }) => {
	const halfWidth = context.board.width / 2.0;
	const centreWidth = Math.floor(halfWidth);
	const leftEdge = Math.ceil(halfWidth) / 2;

	const halfHeight = context.board.height / 2.0;
	const centreHeight = Math.floor(halfHeight);
	const bottomEdge = Math.ceil(halfHeight) / 2;

	const centreCells = [];
	for (let x = leftEdge; x < leftEdge + centreWidth; x += 1) {
		for (let y = bottomEdge; y < bottomEdge + centreHeight; y += 1) {
			const pos = { x, y };
			if (position.IsSafe(pos, context)) {
				centreCells.push(pos);
			}
		}
	}

	if (_.isEmpty(centreCells)) {
		return false;
	}

	const cellIndex = Math.floor(Math.random() * centreCells.length);
	const target = centreCells[cellIndex];
	const move = pathfinding.ApproachTarget(target, context);

	utils.LogMove(context.turn, move, 'GoCentre');
	const isSafe = position.IsSafe(adjacent[move], context);

	return isSafe && move;
};

module.exports = { GoCentre };

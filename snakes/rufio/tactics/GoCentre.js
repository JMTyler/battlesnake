const _ = require('lodash');

const movement = require('../../../movement');
const position = require('../../../position');
const utils    = require('../../../utils');

const GoCentre = () => {
	return ({ context, adjacent }) => {
		const centreWidth = 3;
		const leftEdge = (context.board.width - centreWidth) / 2;

		const centreHeight = 3;
		const bottomEdge = (context.board.height - centreHeight) / 2;

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

		const target = _.sample(centreCells);
		const move = movement.ApproachTarget(target, context);
		utils.LogMove(context.turn, move, 'Go Centre');
		const isSafe = position.IsSafe(adjacent[move], context);

		return isSafe && move;
	};
};

module.exports = { GoCentre };

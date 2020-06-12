const _ = require('lodash');
const position = require('./position');

const GetDistance = (origin, target) => {
	if (_.isArray(target)) {
		return _.map(target, (t) => GetDistance(origin, t));
	}

	const x = target.x - origin.x;
	const y = target.y - origin.y;
	return Math.hypot(x, y);
};

const GetVector = (origin, target) => {
	const x = target.x - origin.x;
	const y = target.y - origin.y;
	return {
		dir: {
			x: Math.sign(x) > 0 ? 'right' : 'left',
			y: Math.sign(y) > 0 ? 'up' : 'down',
		},
		weight: { x, y },
	};
};

const ApproachTarget = (target, { board, you }) => {
	const vector = GetVector(you.head, target);
	const adjacent = position.GetAdjacentTiles(you.head);

	// TODO: Support target being a straight line away, making left/right or up/down equal choices.
	const moveX = adjacent[vector.dir.x];
	if (!position.IsSafe(moveX, { board, you })) {
		return vector.dir.y;
	}

	const moveY = adjacent[vector.dir.y];
	if (!position.IsSafe(moveY, { board, you })) {
		return vector.dir.x;
	}

	if (you.head.x != target.x) {
		return vector.dir.x;
	}

	return vector.dir.y;
};

module.exports = {
	GetDistance,
	GetVector,
	ApproachTarget,
};

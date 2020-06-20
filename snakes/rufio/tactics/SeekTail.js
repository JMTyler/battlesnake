const _ = require('lodash');

const movement = require('../../../movement');

const SeekTail = () => {
	return ({ context }) => {
		return movement.ApproachTarget(_.last(context.you.body), context);
	};
};

module.exports = { SeekTail };

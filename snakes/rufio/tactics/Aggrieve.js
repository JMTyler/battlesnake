const _ = require('lodash');

const movement = require('../../../movement');
const position = require('../../../position');
const utils    = require('../../../utils');

const chooseAdjacentCell = (prey, context, state) => {
	const targetOptions = _.pickBy(position.GetAdjacentTiles(prey.head), (pos) => position.IsSafe(pos, context));
	if (_.isEmpty(targetOptions)) {
		return false;
	}

	const preysLastMove = state.snakes[prey.id].move;
	if (targetOptions[preysLastMove]) {
		return targetOptions[preysLastMove];
	}

	return movement.FindClosestTarget(context.you.head, _.toArray(targetOptions));
};

const Aggrieve = ({ advantage = 1 }) => {
	return ({ context, state }) => {
		const preyOptions = _.filter(context.board.snakes, (snake) => (context.you.length >= snake.length + advantage));
		if (_.isEmpty(preyOptions)) {
			return false;
		}

		const closestPrey = movement.FindClosestTarget(context.you.head, _.map(preyOptions, 'head'));
		const prey = _.find(context.board.snakes, { head: closestPrey });
		const target = chooseAdjacentCell(prey, context, state);
		if (!target) {
			return false;
		}

		const move = movement.ApproachTarget(target, context);
		utils.LogMove(context.turn, move, 'Aggrieve');
		return move;
	};
};

module.exports = { Aggrieve };

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

const EasyKill = ({ advantage = 1, distance = Infinity }) => {
	return ({ context, state }) => {
		// TODO: We should probably just remove `you` from the snakes array at the top level.
		const snakes = _.filter(context.board.snakes, (snake) => snake.id !== context.you.id);
		const closestSnake = movement.FindClosestTarget(context.you.head, _.map(snakes, 'head'));
		const distanceToSnake = movement.GetDistance(context.you.head, closestSnake);
		if (distanceToSnake > distance) {
			return false;
		}
	
		const prey = _.find(snakes, { head: closestSnake });
		if (context.you.length < prey.length + advantage) {
			return false;
		}

		const target = chooseAdjacentCell(prey, context, state);
		if (!target) {
			return false;
		}

		const move = movement.ApproachTarget(target, context);
		utils.LogMove(context.turn, move, 'Easy Kill');
		return move;
	};
};

module.exports = { EasyKill };

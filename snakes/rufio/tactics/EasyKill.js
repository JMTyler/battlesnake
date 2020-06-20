const _ = require('lodash');

const movement = require('../../../movement');
const position = require('../../../position');
const utils    = require('../../../utils');

const EasyKill = ({ advantage = 1, distance = Infinity }) => {
	return ({ context, adjacent }) => {
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

		const targetOptions = _.filter(position.GetAdjacentTiles(prey.head), (pos) => position.IsSafe(pos, context));
		if (_.isEmpty(targetOptions)) {
			return false;
		}

		const target = movement.FindClosestTarget(context.you.head, targetOptions);
		const move = movement.ApproachTarget(target, context);
		utils.LogMove(context.turn, move, 'Easy Kill');
		const isSafe = position.IsSafe(adjacent[move], context);
		return isSafe && move;
	};
};

module.exports = { EasyKill };

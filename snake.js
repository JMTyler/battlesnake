const _ = require('lodash');

const board       = require('./board');
const pathfinding = require('./pathfinding');
const position    = require('./position');
const utils       = require('./utils');

const State = {
	Get(context) {
		return _.get(this, [context.game.id, context.you.id]);
	},
	Set(context, value) {
		return _.set(this, [context.game.id, context.you.id], value);
	},
};

const GetInfo = () => {
	return {
		apiversion: '1',
		author:     'JMTyler',
		color:      '#700070',
		head:       'beluga',
		tail:       'shac-coffee',
	};
};

const Move = (context) => {
	const { you } = context;
	const state = State.Get(context);
	const adjacent = position.GetAdjacentTiles(you.head);

	if (you.health <= state.maxTravel) {
		state.move = pathfinding.ApproachTarget(board.FindClosestFood(context), context);
		utils.LogMove(context.turn, state.move, 'hungry, seeking food');
		return state.move;
	}

	if (position.IsSafe(adjacent[state.move], context)) {
		if (context.turn === 0) utils.LogMove(context.turn, state.move, 'first turn');
		//else Utils.LogMove(context.turn, state.move, 'no change');
		return state.move;
	}

	state.move = pathfinding.ApproachTarget(_.last(you.body), context);
	utils.LogMove(context.turn, state.move, 'unsafe, approaching tail');
	if (position.IsSafe(adjacent[state.move], context)) {
		return state.move;
	}

	state.turns = 0;
	const turn = { right: 'up', up: 'left', left: 'down', down: 'right' };
	do {
		state.move = turn[state.move];
		state.turns += 1;
		utils.LogMove(context.turn, state.move, 'still unsafe, had to turn');
	} while (state.turns < 4 && !position.IsSafe(adjacent[state.move], context));

	if (state.turns === 4) {
		utils.LogMove(context.turn, state.move, 'welp 👋');
	}

	return state.move;
};

const StartGame = (context) => {
	const { board } = context;

	State.Set(context, {
		maxTravel: board.width + board.height - 2,
		turns:     0,
		move:      'right',
	});
	
	console.log('-----');
	console.log();
};

const EndGame = (context) => {
	console.log();
	console.log('* Game Over! *');
};

module.exports = {
	GetInfo,
	StartGame,
	Move,
	EndGame,
};

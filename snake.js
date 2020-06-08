const _ = require('lodash');

const tactics = require('./tactics');
const utils   = require('./utils');

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
	
	let isSafe, move;

	if (you.health <= state.maxTravel) {
		[isSafe, move] = tactics.SeekFood(context);
		state.move = move;
		return move;
	}

	[isSafe, move] = tactics.Continue(state.move, context);
	state.move = move;
	if (isSafe) {
		return move;
	}

	[isSafe, move] = tactics.SeekTail(context);
	state.move = move;
	if (isSafe) {
		return move;
	}

	[isSafe, move] = tactics.RotateUntilSafe(state.move, context);
	state.move = move;
	if (isSafe) {
		return move;
	}

	utils.LogMove(context.turn, move, 'welp 👋');
	return move;
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

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
	const state = State.Get(context);
	const adjacent = position.GetAdjacentTiles(context.you.head);

	let move;

	move = tactics.SeekFood({ context, state, adjacent });
	if (move) {
		state.move = move;
		return move;
	}

	move = tactics.Continue({ context, state, adjacent });
	if (move) {
		state.move = move;
		return move;
	}

	move = tactics.SeekTail({ context, state, adjacent });
	if (move) {
		state.move = move;
		return move;
	}

	move = tactics.RotateUntilSafe({ context, state, adjacent });
	if (move) {
		state.move = move;
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

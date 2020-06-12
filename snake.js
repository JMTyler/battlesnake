const _ = require('lodash');

const position = require('./position');
const tactics  = require('./tactics');
const utils    = require('./utils');

const State = {
	Initialise(context, value) {
		return _.set(this, [context.game.id, context.you.id], value);
	},
	Scope(context) {
		return _.get(this, [context.game.id, context.you.id]);
	},
};

const GetInfo = () => {
	return {
		apiversion: '1',
		author:     'JMTyler',
		color:      '#700070',
		head:       'pixel',
		tail:       'pixel',
	};
};

const strategy = [
	tactics.Aggress,
	tactics.SeekFood,
	tactics.Continue,
	tactics.SeekTail,
	tactics.RotateUntilSafe,
];

const Move = (context) => {
	const state = State.Scope(context);
	const adjacent = position.GetAdjacentTiles(context.you.head);
	
	const move = _.reduce(strategy, (prev, tactic) => {
		return prev || tactic({ context, state, adjacent });
	}, false);
	
	if (move) {
		state.move = move;
		return move;
	}

	utils.LogMove(context.turn, state.move, 'welp 👋');
	return state.move;
};

const StartGame = (context) => {
	State.Initialise(context, {
		move: 'right',
	});
	
	console.log('-----');
	console.log();
};

const EndGame = (context) => {
	const result = (context.you.id === _.get(context, 'board.snakes.0.id')) ? 'WIN' : 'LOSE';
	console.log();
	console.log(`* Game Over! ${result} *`);
};

module.exports = {
	GetInfo,
	StartGame,
	Move,
	EndGame,
};

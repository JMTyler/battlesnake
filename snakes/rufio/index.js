const _ = require('lodash');

const movement = require('../../movement');
const position = require('../../position');
const utils    = require('../../utils');

const State = {
	Initialise(context, value) {
		return _.set(this, [context.game.id, context.you.id], value);
	},
	Scope(context) {
		return _.get(this, [context.game.id, context.you.id], { move: 'right' });
	},
};

const GetInfo = () => {
	return {
		color : '#8F008F',
		head  : 'shades',
		tail  : 'bolt',
	};
};

const tactics = require('./tactics');
const strategy = [
	tactics.EasyKill({ advantage: 1, distance: 2 }),
	tactics.EasySnack({ distance: 2 }),
	tactics.Abscond({ disadvantage: 1, distance: 3 }),
	tactics.Aggrieve({ advantage: 2 }),
	tactics.Hungry({ health: 95 }),
	tactics.GoCentre(),
	tactics.Continue(),
	tactics.SeekTail(),
	tactics.RotateUntilSafe(),
];

const Move = async (context) => {
	await utils.RecordFrame(context);

	const state = State.Scope(context);
	const adjacent = position.GetAdjacentTiles(context.you.head);

	movement.InitPathfinder(context);

	const move = _.reduce(strategy, (prev, tactic) => {
		return prev || tactic({ context, state, adjacent });
	}, false);

	await utils.RecordFrame(context, move || state.move);

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

const EndGame = async (context) => {
	const result = (context.you.id === _.get(context, 'board.snakes.0.id')) ? 'WIN' : 'LOSE';
	console.log();
	console.log(`* Game Over! ${result} *`);

	await utils.PruneGames();
};

module.exports = {
	GetInfo,
	StartGame,
	Move,
	EndGame,
};

const _ = require('lodash');

const movement = require('../../movement');
const position = require('../../position');
const utils    = require('../../utils');

const State = {
	Initialise(context, value) {
		return _.set(this, [context.game.id, context.you.id], value);
	},
	Scope(context) {
		return _.get(this, [context.game.id, context.you.id], { move: 'right', snakes: {} });
	},
};

const tactics = require('./tactics');
const strategy = [
	tactics.EasyKill({ advantage: 1, distance: 2 }),
	tactics.EasySnack({ distance: 2 }),
	tactics.Abscond({ disadvantage: 1, distance: 3 }),
	tactics.Aggrieve({ advantage: 2 }),
	tactics.Hungry({}),
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

	// Figure out which move each snake took during the *last* turn.
	state.snakes = _.chain(context.board.snakes)
		.mapKeys('id')
		.mapValues(({ id, head }) => {
			const prev = _.get(state.snakes, [id, 'head']);
			const move = prev ? position.ToDirection(head, prev) : 'up';
			return { head, move };
		})
		.value();

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

const GetInfo = () => require('./info.json');

const StartGame = async (context) => {
	State.Initialise(context, {
		move   : 'right',
		snakes : {},
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

module.exports = { GetInfo, StartGame, Move, EndGame };

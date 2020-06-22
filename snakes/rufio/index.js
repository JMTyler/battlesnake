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
const strategy = {
	'Easy Kill'         : tactics.Aggrieve({ advantage: 1, distance: 1 }),
	'Quick Snack'       : tactics.Eat({ distance: 2 }),
	'Abscond'           : tactics.Abscond({ disadvantage: 1, distance: 3 }),
	'Aggrieve'          : tactics.Aggrieve({ advantage: 2 }),
	'Hungry'            : tactics.Eat(),
	'Go Centre'         : tactics.GoCentre(),
	'Continue'          : tactics.Continue(),
	'Seek Tail'         : tactics.SeekTail(),
	'Rotate Until Safe' : tactics.RotateUntilSafe(),
};

const Move = async (context) => {
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

	_.remove(context.board.snakes, context.you);
	const move = _.reduce(strategy, (prev, tactic, description) => {
		if (prev) return prev;
		const move = tactic({ context, state, adjacent });
		if (!move) return false;
		utils.LogMove(context.turn, move, description);
		const isSafe = position.IsSafe(adjacent[move], context);
		return isSafe && move;
	}, false);

	if (move) {
		state.move = move;
		return move;
	}

	utils.LogMove(context.turn, state.move, 'welp 👋');
	return state.move;
};

const GetInfo = async () => require('./info.json');

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

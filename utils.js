const _ = require('lodash');

const config = require('./config');

let db = require('./db');
db.Connect(config.get('database_url'))
	.then((_db) => {
		db = _db;
	});

const Leftpad = (string, pad = 5) => {
	string = _.toString(string);
	const prefix = _.times(pad - string.length, () => ' ').join('');
	return prefix + string;
};

let previousTurn = null;
const LogMove = (turn, move, comment) => {
	if (turn > previousTurn + 1) {
		console.log(' [ ... ]');
	}

	const moveTag = Leftpad(move);
	let turnTag = `[${Leftpad(turn)}]`;
	if (turn === previousTurn) {
		turnTag = ` ${Leftpad('↳')} `;
	}

	previousTurn = turn;

	console.log(` ${turnTag} ${moveTag} :  ${comment}`);
};

// TODO: Merge LogMove and RecordFrame, once we can be sure which move was the final choice.
const RecordFrame = async (context, move = null) => {
	const NOW = new Date().toISOString();

	if (move) {
		return await db.Frames.update({
			game_id  : context.game.id,
			snake_id : context.you.id,
			turn     : context.turn,
		}, {
			move,
			updated_at : NOW,
		});
	}

	return await db.Frames.insert({
		context,
		game_id    : context.game.id,
		snake_id   : context.you.id,
		name       : context.you.name,
		turn       : context.turn,
		created_at : NOW,
		updated_at : NOW,
	}, { onConflict: { action: 'ignore' } });
};

const PruneGames = async () => {
	const numRows = _.toNumber(await db.Frames.count());
	if (numRows < 9000) return;

	// Find the oldest game in the database.
	const { game_id } = await db.Frames.findOne({}, {
		fields : ['game_id'],
		order  : [{ field: 'created_at', direction: 'asc' }],
	});

	// And delete it.
	return await db.Frames.destroy({ game_id });
};

module.exports = {
	Leftpad,
	LogMove,
	RecordFrame,
	PruneGames,
};

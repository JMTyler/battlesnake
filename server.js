const _ = require('lodash');
const express = require('express');

const app = express();
app.use(require('body-parser').json());

app.get('/', (req, res) => {
	res.send({
		apiversion: '1',
		author:     'JMTyler',
		color:      '#700070',
		head:       'beluga',
		tail:       'shac-coffee',
	});
});

const State = {
	Get(context) {
		return _.get(this, [context.game.id, context.you.id]);
	},
	Set(context, value) {
		return _.set(this, [context.game.id, context.you.id], value);
	},
};

const Utils = {
	Leftpad(string, pad = 5) {
		string = _.toString(string);
		const prefix = _.times(pad - string.length, () => ' ').join('');
		return prefix + string;
	},

	previousTurn: null,
	LogMove(turn, move, comment) {
		if (turn > this.previousTurn + 1) {
			console.log(' [ ... ]');
		}
		
		const moveTag = Utils.Leftpad(move);
		let turnTag = `[${Utils.Leftpad(turn)}]`;
		if (turn === this.previousTurn) {
			turnTag = ` ${Utils.Leftpad('↳')} `;
		}

		this.previousTurn = turn;

		console.log(` ${turnTag} ${moveTag} :  ${comment}`);
	},
};

const Path = {
	ApproachTarget(target, { board, you }) {
		const vector = Path.GetVector(you.head, target);
		const adjacent = Position.GetAdjacentTiles(you.head);

		const moveX = adjacent[vector.dir.x];
		if (!Position.IsSafe(moveX, { board, you })) {
			return vector.dir.y;
		}

		const moveY = adjacent[vector.dir.y];
		if (!Position.IsSafe(moveY, { board, you })) {
			return vector.dir.x;
		}

		if (you.head.x != target.x) {
			return vector.dir.x;
		}

		return vector.dir.y;
	},

	GetVector(origin, dest) {
		const x = dest.x - origin.x;
		const y = dest.y - origin.y;
		return {
			dir: {
				x: Math.sign(x) < 0 ? 'left' : 'right',
				y: Math.sign(y) < 0 ? 'down' : 'up',
			},
			weight: { x, y },
		};
	},

	GetDistance(origin, target) {
		if (_.isArray(target)) {
			return _.map(target, (t) => Path.GetDistance(origin, t));
		}

		const x = target.x - origin.x;
		const y = target.y - origin.y;
		return Math.hypot(x, y);
	},
};

const Position = {
	GetAdjacentTiles({ x, y }) {
		return {
			up:    { x, y: y + 1 },
			down:  { x, y: y - 1 },
			left:  { x: x - 1, y },
			right: { x: x + 1, y },
		};
	},

	IsOutsideBoard({ x, y }, board) {
		return x < 0 || y < 0 || x >= board.width || y >= board.height;
	},

	IsDeadly(pos, { board }) {
		if (Position.IsOutsideBoard(pos, board)) {
			return true;
		}

		const anySnakeCollision = _.some(board.snakes, (snake) => {
			// TODO: Only drop the tail piece if the snake HASN'T just eaten a disc.
			// TODO: Or, drop the tail either way, but consider that spot risky.
			return _.some(_.initial(snake.body), pos);
		});
		if (anySnakeCollision) {
			return true;
		}

		return false;
	},

	IsRisky(pos, { board, you }) {
		return _.some(board.snakes, (snake) => {
			if (snake.id === you.id) return false;
			const adjacent = Position.GetAdjacentTiles(snake.head);
			return _.some(adjacent, pos);
		});
	},

	IsSafe(pos, context) {
		return !Position.IsDeadly(pos, context) && !Position.IsRisky(pos, context);
	},
};

const Food = {
	FindClosest({ you, board }) {
		const distances = Path.GetDistance(you.head, board.food);
		const shortestIndex = _.reduce(distances, (prev, distance, ix) => {
			return (distance < distances[prev]) ? ix : prev;
		}, 0);
		return board.food[shortestIndex];
	},
};

app.post('/start', (req, res) => {
	const { board } = req.body;

	State.Set(req.body, {
		maxTravel: board.width + board.height - 2,
		turns:     0,
		target:    null,
		move:      'right',
	});
	
	console.log('-----');
	console.log();

	return res.sendStatus(200);
});

app.post('/move', (req, res) => {
	const context = req.body;
	const { board, you } = context;
	const state = State.Get(context);
	const adjacent = Position.GetAdjacentTiles(you.head);

	if (_.isEqual(you.head, state.target)) state.target = null;
	if (state.target) {
		state.move = Path.ApproachTarget(state.target, context);
		Utils.LogMove(context.turn, state.move, 'continuing to approach target');
		return res.send({ move: state.move });
	}

	if (you.health <= state.maxTravel) {
		state.move = Path.ApproachTarget(Food.FindClosest({ you, board }), context);
		Utils.LogMove(context.turn, state.move, 'hungry, seeking food');
		return res.send({ move: state.move });
	}

	if (Position.IsSafe(adjacent[state.move], { board, you })) {
		if (context.turn === 0) Utils.LogMove(context.turn, state.move, 'first turn');
		//else Utils.LogMove(context.turn, state.move, 'no change');
		return res.send({ move: state.move });
	}

	state.move = Path.ApproachTarget(_.last(you.body), context);
	Utils.LogMove(context.turn, state.move, 'unsafe, approaching tail');
	if (Position.IsSafe(adjacent[state.move], { board, you })) {
		return res.send({ move: state.move });
	}

	state.turns = 0;
	const turn = { right: 'up', up: 'left', left: 'down', down: 'right' };
	do {
		state.move = turn[state.move];
		state.turns += 1;
		Utils.LogMove(context.turn, state.move, 'still unsafe, had to turn');
	} while (state.turns < 4 && !Position.IsSafe(adjacent[state.move], { board, you }));

	if (state.turns === 4) {
		Utils.LogMove(context.turn, state.move, 'welp 👋');
	}

	return res.send({ move: state.move });
});

app.post('/end', (req, res) => {
	console.log();
	console.log('* Game Over! *');
	return res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!'));

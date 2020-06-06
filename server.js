const _       = require('lodash');
const express = require('express');

const board       = require('./board');
const pathfinding = require('./pathfinding');
const position    = require('./position');
const utils       = require('./utils');

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
	const { you } = context;
	const state = State.Get(context);
	const adjacent = position.GetAdjacentTiles(you.head);

	if (_.isEqual(you.head, state.target)) state.target = null;
	if (state.target) {
		state.move = pathfinding.ApproachTarget(state.target, context);
		utils.LogMove(context.turn, state.move, 'continuing to approach target');
		return res.send({ move: state.move });
	}

	if (you.health <= state.maxTravel) {
		state.move = pathfinding.ApproachTarget(board.FindClosestFood(context), context);
		utils.LogMove(context.turn, state.move, 'hungry, seeking food');
		return res.send({ move: state.move });
	}

	if (position.IsSafe(adjacent[state.move], context)) {
		if (context.turn === 0) utils.LogMove(context.turn, state.move, 'first turn');
		//else Utils.LogMove(context.turn, state.move, 'no change');
		return res.send({ move: state.move });
	}

	state.move = pathfinding.ApproachTarget(_.last(you.body), context);
	utils.LogMove(context.turn, state.move, 'unsafe, approaching tail');
	if (position.IsSafe(adjacent[state.move], context)) {
		return res.send({ move: state.move });
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

	return res.send({ move: state.move });
});

app.post('/end', (req, res) => {
	console.log();
	console.log('* Game Over! *');
	return res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!'));

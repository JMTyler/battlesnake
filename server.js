const _ = require('lodash');
const express = require('express');

const app = express();
app.use(require('body-parser').json());

app.get('/', (req, res) => {
	res.send({
		apiversion: '1',
		author: 'JMTyler',
		color: '#700070',
		head: 'beluga',
		tail: 'shac-coffee',
	});
});

const Path = {
	ApproachTarget(you, target) {
		console.log('You:', you.head);
		console.log('Target:', target);
		
		const vector = Path.GetVector(you.head, target);
		const adjacent = Position.GetAdjacentTiles(you.head);
		
		const moveX = adjacent[vector.dir.x];
		if (Position.IsDeadly(moveX, { you })) {
			return vector.dir.y;
		}
		
		const moveY = adjacent[vector.dir.y];
		if (Position.IsDeadly(moveY, { you })) {
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
		console.log('DiffX:', x, ';  DiffY:', y);
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
	
	IsDeadly(pos, { board, you }) {
		if (board && Position.IsOutsideBoard(pos, board)) {
			return true;
		}
		
		// TODO: Only drop the tail piece if we HAVEN'T just eaten a disc.
		const collides = _.some(_.initial(you.body), pos);
		if (collides) {
			return true;
		}
		
		return false;
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

let maxTravel = 0;
app.post('/start', (req, res) => {
	const { game, board, you } = req.body;
	
	console.log(`New Game! [${game.id}]`);
	maxTravel = board.width + board.height - 2;
	
	return res.sendStatus(200);
});

let turns = 0;
let target = null;
let move = 'right';
app.post('/move', (req, res) => {
	const { game, board, you } = req.body;
	const adjacent = Position.GetAdjacentTiles(you.head);
	turns = 0;
	
	if (_.isEqual(you.head, target)) target = null;
	if (target) {
		move = Path.ApproachTarget(you, target);
		return res.send({ move });
	}
	
	if (you.health <= maxTravel) {
		move = Path.ApproachTarget(you, Food.FindClosest({ you, board }));
		return res.send({ move });
	}
	
	const turn = { right: 'up', up: 'left', left: 'down', down: 'right' };
	while (turns < 3 && Position.IsDeadly(adjacent[move], { board, you })) {
		move = turn[move];
		turns += 1;
	}
	
	return res.send({ move });
});

app.post('/end', (req, res) => {
	console.log('GAMEOVER', req.body);
	console.log('-----\n');
	return res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!\n-----\n'));

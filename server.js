const express = require('express');

const app = express();
app.use(require('body-parser').json());

app.get('/', (req, res) => {
	res.send({
		apiversion: '1',
		author: 'JMTyler',
		color: '#660066',
		head: 'default',
		tail: 'default',
	});
});

let target = null;
const approachTarget = (you, target) => {
	console.log('You:', you.head);
	console.log('Target:', target);
	
	const diffX = target.x - you.head.x;
	const diffY = target.y - you.head.y;
	
	const absX = Math.abs(diffX);
	const absY = Math.abs(diffY);
	
	console.log('DiffX:', diffX, ';  DiffY:', diffY);
	const moveX = Math.sign(diffX) + you.head.x;
	const moveY = Math.sign(diffY) + you.head.y;
	if (moveX != you.body[1].x) {
		if (absX >= absY || moveY == you.body[1].y) {
			const move = Math.sign(diffX) < 0 ? 'left' : 'right';
			console.log('Moving X:', move);
			return move;
		}
	}
	
	const move = Math.sign(diffY) < 0 ? 'down' : 'up';
	console.log('Moving Y:', move);
	return move;
};

let maxTravel = 0;
app.post('/start', (req, res) => {
	const { game, board, you } = req.body;
	
	console.log(`New Game! [${game.id}]`);
	maxTravel = board.width + board.height - 2;
	
	return res.sendStatus(200);
});

let move = 'right';
app.post('/move', (req, res) => {
	const { game, board, you } = req.body;
	
	if (target) {
		if (you.head.x == target.x && you.head.y == target.y) target = null;
		else {
			move = approachTarget(you, target);
			return res.send({ move });
		}
	}
	
	if (you.health <= maxTravel) {
		target = board.food[0];
		move = approachTarget(you, target);
		return res.send({ move });
	}
	
	const engine = {
		right: { next: 'up',    limit: (you.head.x + 1 >= board.width) },
		up:    { next: 'left',  limit: (you.head.y + 1 >= board.height) },
		left:  { next: 'down',  limit: (you.head.x - 1 < 0) },
		down:  { next: 'right', limit: (you.head.y - 1 < 0) },
	};
	
	if (engine[move].limit) move = engine[move].next;
	
	let x = you.head.x;
	let y = you.head.y;
	if (move == 'up') y += 1;
	if (move == 'down') y -= 1;
	if (move == 'right') x += 1;
	if (move == 'left') x -= 1;
	
	const collides = you.body.reduce((hit, coords) => {
		return hit || (x == coords.x && y == coords.y);
	}, false);
	
	if (collides) {
		target = you.body[you.body.length - 1];
		move = approachTarget(you, target);
		return res.send({ move });
	}
	
	return res.send({ move });
});

app.post('/end', (req, res) => {
	console.log('GAMEOVER', req.body);
	console.log('-----\n');
	return res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!'));

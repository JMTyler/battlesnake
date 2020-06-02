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
	
	if (you.health <= maxTravel) {
		const target = board.food[0];
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
				move = Math.sign(diffX) < 0 ? 'left' : 'right';
				console.log('Moving X:', move);
				return res.send({ move });
			}
		}
		
		move = Math.sign(diffY) < 0 ? 'down' : 'up';
		console.log('Moving Y:', move);
		return res.send({ move });
	}
	
	const engine = {
		right: { next: 'up',    limit: (you.head.x + 1 >= board.width) },
		up:    { next: 'left',  limit: (you.head.y + 1 >= board.height) },
		left:  { next: 'down',  limit: (you.head.x - 1 < 0) },
		down:  { next: 'right', limit: (you.head.y - 1 < 0) },
	};
	
	if (engine[move].limit) move = engine[move].next;
	
	return res.send({ move });
});

app.post('/end', (req, res) => {
	console.log('GAMEOVER', req.body);
	console.log('-----\n');
	return res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!'));

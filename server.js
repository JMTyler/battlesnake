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

app.post('/start', (req, res) => {
	console.log(`New Game! [${req.body.game.id}]`)
	res.sendStatus(200);
});

let move = 'right';
app.post('/move', (req, res) => {
	const { game, board, you } = req.body;
	
	const engine = {
		right: { next: 'up',    limit: (you.head.x + 1 >= board.width) },
		up:    { next: 'left',  limit: (you.head.y + 1 >= board.height) },
		left:  { next: 'down',  limit: (you.head.x - 1 < 0) },
		down:  { next: 'right', limit: (you.head.y - 1 < 0) },
	};
	
	if (engine[move].limit) move = engine[move].next;
	
	res.send({ move });
});

app.post('/end', (req, res) => {
	console.log('GAMEOVER', req.body);
	console.log('-----\n');
	res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!'));

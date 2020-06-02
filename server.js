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

app.post('/move', (req, res) => {
	const { game, board, you } = req.body;
	
	let move = 'right';
	if (you.head.x + 1 >= board.width) move = 'up';
	if (you.head.y + 1 >= board.height) move = 'left';
	if (you.head.x - 1 < 0) move = 'down';
	if (you.head.y - 1 < 0) move = 'right';
	
	res.send({ move });
});

app.post('/end', (req, res) => {
	console.log('GAMEOVER', req.body);
	console.log('-----\n');
	res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!'));

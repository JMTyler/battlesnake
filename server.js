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
	res.send({ move: 'right' });
});

app.post('/end', (req, res) => {
	console.log('GAMEOVER', req.body);
	console.log('-----\n');
	res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => console.log('Running!'));

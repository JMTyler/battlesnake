const _       = require('lodash');
const express = require('express');

const snake = require('./snake');

const app = express();
app.use(require('body-parser').json());

app.get('/', (req, res) => {
	res.send(snake.GetInfo());
});

app.post('/start', (req, res) => {
	snake.StartGame(req.body);
	return res.sendStatus(200);
});

app.post('/move', (req, res) => {
	const move = snake.Move(req.body);
	return res.send({ move });
});

app.post('/end', (req, res) => {
	snake.EndGame(req.body);
	return res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => {
	console.log('Running!');
});

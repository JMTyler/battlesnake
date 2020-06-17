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

app.post('/move', async (req, res) => {
	const start = Date.now();
	const move = await snake.Move(req.body);
	console.log(`Move took ${Date.now() - start}ms.`);
	return res.send({ move });
});

app.post('/end', async (req, res) => {
	await snake.EndGame(req.body);
	return res.sendStatus(200);
});

app.listen(process.env.PORT || 9000, () => {
	console.log('Running!');
});

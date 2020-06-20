const _       = require('lodash');
const express = require('express');
const fs      = require('fs');

const files = fs.readdirSync('./snakes');

// TODO: Setup root paths to default to local snake.
module.exports = _.transform(files, (router, file) => {
	const name = file.replace(/\.js$/, '');
	const snake = require(`./snakes/${file}`);

	router.get(`/${name}`, (req, res) => {
		res.send(_.extend({
			apiversion : '1',
			author     : 'JMTyler',
		}, snake.GetInfo()));
	});

	router.post(`/${name}/start`, async (req, res) => {
		await snake.StartGame(req.body);
		return res.sendStatus(200);
	});

	router.post(`/${name}/move`, async (req, res) => {
		const start = Date.now();
		const move = await snake.Move(req.body);
		console.log(`Move took ${Date.now() - start}ms.`);
		return res.send({ move });
	});

	router.post(`/${name}/end`, async (req, res) => {
		await snake.EndGame(req.body);
		return res.sendStatus(200);
	});
}, express.Router());

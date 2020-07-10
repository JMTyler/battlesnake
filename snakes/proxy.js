const request = require('request');

const baseUrl = 'http://battlesnake.jaredtyler.ca/rufio';

const GetInfo = async () => {
	return new Promise((resolve, reject) => {
		request.get({
			url: '/',
			baseUrl,
			json: true,
		}, (err, res, body) => err ? reject(err) : resolve(body));
	});
};

const StartGame = async (context) => {
	return new Promise((resolve, reject) => {
		request.post({
			url: '/start',
			body: context,
			baseUrl,
			json: true,
		}, (err) => err ? reject(err) : resolve());
	});
};

const Move = async (context) => {
	return new Promise((resolve, reject) => {
		request.post({
			url: '/move',
			body: context,
			baseUrl,
			json: true,
		}, (err, res, body) => err ? reject(err) : resolve(body.move));
	});
};

const EndGame = async (context) => {
	return new Promise((resolve, reject) => {
		request.post({
			url: '/end',
			body: context,
			baseUrl,
			json: true,
		}, (err) => err ? reject(err) : resolve());
	});
};

module.exports = { GetInfo, StartGame, Move, EndGame };

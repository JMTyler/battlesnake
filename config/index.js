module.exports = require('pico-conf')
	.argv()
	.env()
	.file(`./${process.env.NODE_ENV}.js`, { silent: true })
	.file('./default.js');

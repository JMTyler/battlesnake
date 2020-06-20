const express = require('express');

const config = require('./config');

const app = express();
app.use(require('body-parser').json());
app.use(require('./router'));
app.listen(config.get('port'), () => {
	console.log('Running!');
});

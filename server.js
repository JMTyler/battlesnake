const express = require('express');

const app = express();

app.get('/', (req, res) => {
	console.log('bleep');
	res.send('bloop');
});

app.listen(process.env.PORT || 9000);

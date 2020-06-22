const { StartGame, Move, EndGame } = require('./rufio');

const GetInfo = async () => {
	return {
		color : '#008F00',
		head  : 'shac-workout',
		tail  : 'freckled',
	};
};

module.exports = { GetInfo, StartGame, Move, EndGame };

const _ = require('lodash');

const Leftpad = (string, pad = 5) => {
	string = _.toString(string);
	const prefix = _.times(pad - string.length, () => ' ').join('');
	return prefix + string;
};

let previousTurn = null;
const LogMove = (turn, move, comment) => {
	if (turn > previousTurn + 1) {
		console.log(' [ ... ]');
	}
	
	const moveTag = Leftpad(move);
	let turnTag = `[${Leftpad(turn)}]`;
	if (turn === previousTurn) {
		turnTag = ` ${Leftpad('↳')} `;
	}

	previousTurn = turn;

	console.log(` ${turnTag} ${moveTag} :  ${comment}`);
};

module.exports = {
	Leftpad,
	LogMove,
};

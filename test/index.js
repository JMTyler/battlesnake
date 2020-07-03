const _ = require('lodash');

const engine = require('../../../stolksdorf/battlesnake/lib/snek.engine');
const utils  = require('../../../stolksdorf/battlesnake/lib/utils');

const snakes = {
	'Local Dev'           : require('../snakes/local'),
	'Rufio the Tenacious' : require('../snakes/proxy'),
};

const states = require('./game-state.json');

let foodSpawns;
let initial = states;
if (_.isArray(states)) {
	initial = states[0];
	foodSpawns = _.chain(states)
		.map(({ turn, board: { food } }, ix) => {
			if (ix > 0) {
				food = _.differenceWith(food, states[ix-1].board.food, _.isEqual);
			}
			return { turn, food };
		})
		.reject(({ food }) => _.isEmpty(food))
		.mapKeys('turn')
		.mapValues('food')
		.value();
}

initial.game.dev = true;
engine.play(snakes, initial, foodSpawns, (state) => {
	console.log(utils.print(state) + '\n');
})
.catch(console.error)
.then(() => process.exit(0));

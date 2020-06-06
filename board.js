const _ = require('lodash');
const pathfinding = require('./pathfinding');

const FindClosestFood = ({ you, board }) => {
	const distances = pathfinding.GetDistance(you.head, board.food);
	const shortestIndex = _.reduce(distances, (prev, distance, ix) => {
		return (distance < distances[prev]) ? ix : prev;
	}, 0);
	return board.food[shortestIndex];
};

module.exports = {
	FindClosestFood,
};

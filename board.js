const _ = require('lodash');
const pathfinding = require('./pathfinding');

const FindClosestFood = ({ you, board }) => {
	return pathfinding.FindClosestTarget(you.head, board.food);
};

module.exports = {
	FindClosestFood,
};

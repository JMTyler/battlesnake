const movement = require('./movement');

const FindClosestFood = ({ you, board }) => {
	return movement.FindClosestTarget(you.head, board.food);
};

module.exports = {
	FindClosestFood,
};

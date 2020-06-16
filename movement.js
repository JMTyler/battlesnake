const _           = require('lodash');
const pathfinding = require('pathfinding');

const position = require('./position');

const pathfinder = new pathfinding.JumpPointFinder({ diagonalMovement: pathfinding.DiagonalMovement.Never });

const InitPathfinder = (context) => {
	const grid = new pathfinding.Grid(context.board.width, context.board.height);

	// TODO: Consider adding a safeGrid (this is a riskyGrid) that also avoids risky cells.
	_.each(context.board.snakes, (snake) => {
		// TODO: not sure if this field exists in regular payloads
		if (snake.death) return;
		_.each(_.initial(snake.body), ({ x, y }) => {
			grid.setWalkableAt(x, y, false);
		});
	});

	// Make context.grid an accessor that always returns a clone.
	Object.defineProperty(context, 'grid', { get: () => grid.clone() });
};

const GetDistance = (origin, target) => {
	if (_.isArray(target)) {
		return _.map(target, (t) => GetDistance(origin, t));
	}

	const x = Math.abs(target.x - origin.x);
	const y = Math.abs(target.y - origin.y);
	return x + y;
};

const GetVector = (origin, target) => {
	const x = target.x - origin.x;
	const y = target.y - origin.y;
	return {
		dir: {
			x: Math.sign(x) > 0 ? 'right' : 'left',
			y: Math.sign(y) > 0 ? 'up' : 'down',
		},
		weight: { x, y },
	};
};

const ApproachTarget = (target, { you, grid }) => {
	const path = pathfinder.findPath(you.head.x, you.head.y, target.x, target.y, grid);
	const nextCell = path[1];
	if (!nextCell) return 'up';
	const pos = { x: nextCell[0], y: nextCell[1] };
	return position.ToDirection(pos, { you });
};

const FindClosestTarget = (origin, targets) => {
	const distances = GetDistance(origin, targets);
	const shortestIndex = _.reduce(distances, (prev, distance, ix) => {
		return (distance < distances[prev]) ? ix : prev;
	}, 0);
	return targets[shortestIndex];
};

module.exports = {
	InitPathfinder,
	GetDistance,
	GetVector,
	ApproachTarget,
	FindClosestTarget,
};

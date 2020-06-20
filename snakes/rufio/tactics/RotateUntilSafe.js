const position = require('../../../position');

const RotateUntilSafe = () => {
	return ({ context, state, adjacent }) => {
		const rotate = { right: 'down', down: 'left', left: 'up', up: 'right' };

		let turns = 0;
		let isSafe = false;
		let move = state.move;
		do {
			move = rotate[move];
			turns += 1;
			isSafe = position.IsSafe(adjacent[move], context);
//			utils.LogMove(context.turn, move, 'Rotate Until Safe');
		} while (turns < 4 && !isSafe);

		if (isSafe) {
			return move;
		}

		// If there are no safe cells nearby, we have to be willing to move into risky cells.
		// Prioritising our current direction.
		if (!position.IsDeadly(adjacent[state.move], context)) {
			return state.move;
		}

		// But if our current direction is deadly, resort to any adjacent risky cell.
		turns = 0;
		do {
			move = rotate[move];
			turns += 1;
			isSafe = !position.IsDeadly(adjacent[move], context);
//			utils.LogMove(context.turn, move, 'Rotate Until Risky');
		} while (turns < 4 && !isSafe);

		return move;
	};
};

module.exports = { RotateUntilSafe };

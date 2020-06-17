const massive = require('massive');

let db;
let connected = false;

module.exports = {
	/**
	 * @returns {Promise<massive.Database>}
	 */
	async Connect(connInfo) {
		if (connected) return db;

		db = await massive(connInfo);
		await require('./migrate')(db);

		// This is really weird but it'll do for now.
		module.exports = db;
		connected = true;
		return db;
	},
};

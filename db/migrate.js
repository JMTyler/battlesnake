
module.exports = async (db) => {
	// TODO: Use an actual migration library, not raw queries.
	await db.withTransaction(async (tx) => {
		await tx.query(`
			CREATE TABLE IF NOT EXISTS "Frames" (
				id         bigserial   NOT NULL,
				created_at timestamptz NOT NULL DEFAULT now(),
				updated_at timestamptz NOT NULL DEFAULT now(),

				game_id  uuid    NOT NULL,
				snake_id text    NOT NULL,
				name     text    NOT NULL,
				turn     integer NOT NULL,
				move     text,
				context  jsonb,

				PRIMARY KEY (id),
				UNIQUE (game_id, turn, snake_id)
			);
		`);

		await tx.query(`
			ALTER TABLE "Frames"
			ADD COLUMN IF NOT EXISTS important boolean NOT NULL DEFAULT false;
		`);
	});

	// Tell Massive to introspect the database again now that we've modified the schema.
	await db.reload();
};

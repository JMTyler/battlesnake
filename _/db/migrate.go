package db

import (
	"github.com/go-pg/pg/v9"
)

func migrate() error {
	// TODO: Use an actual migration library, not just raw idempotent queries.
	return DB.RunInTransaction(func(tx *pg.Tx) error {
		if _, err := tx.Exec(`
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
		`); err != nil {
			return err
		}

		if _, err := tx.Exec(`
			ALTER TABLE "Frames"
			ADD COLUMN IF NOT EXISTS important boolean NOT NULL DEFAULT false;
		`); err != nil {
			return err
		}

		if _, err := tx.Exec(`
			ALTER TABLE "Frames"
			ADD COLUMN IF NOT EXISTS duration integer DEFAULT NULL;
		`); err != nil {
			return err
		}

		return nil
	})
}

package migration

func init() {
	sourceDriver.append(
		202401221049,
		`
			CREATE TABLE IF NOT EXISTS carts (
				id char(36) NOT NULL,
				name varchar(255) NULL,
				is_active bool NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT carts_pk PRIMARY KEY (id)
			);
		`,
		`
			DROP TABLE IF EXISTS carts;
		`,
	)
}

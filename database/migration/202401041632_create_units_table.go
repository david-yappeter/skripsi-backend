package migration

func init() {
	sourceDriver.append(
		202401041632,
		`
			CREATE TABLE IF NOT EXISTS units (
				id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT units_pk PRIMARY KEY (id),
				CONSTRAINT units_uk_name UNIQUE (name)
			);
		`,
		`
			DROP TABLE IF EXISTS units;
		`,
	)
}

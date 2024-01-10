package migration

func init() {
	sourceDriver.append(
		202311131520,
		`
			CREATE TABLE IF NOT EXISTS files (
				id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				type varchar(255) NOT NULL,
				path text NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT files_pk PRIMARY KEY (id)
			);
		`,
		`
			DROP TABLE IF EXISTS files;
		`,
	)
}

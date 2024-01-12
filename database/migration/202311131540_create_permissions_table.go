package migration

func init() {
	sourceDriver.append(
		202311131540,
		`
			CREATE TABLE IF NOT EXISTS permissions (
				id char(36) NOT NULL,
				title varchar(100) NOT NULL,
				description varchar(255) NOT NULL,
				is_active bool NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT permissions_pk PRIMARY KEY (id),
				CONSTRAINT permission_title_uk UNIQUE (title)
			);
		`,
		`
			DROP TABLE IF EXISTS permissions;
		`,
	)
}

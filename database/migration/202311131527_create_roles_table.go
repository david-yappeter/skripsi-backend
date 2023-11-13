package migration

func init() {
	sourceDriver.append(
		202311131527,
		`
			CREATE TABLE IF NOT EXISTS roles (
				id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				description text NULL,
				created_at timestamp,
				updated_at timestamp,
				CONSTRAINT roles_pk PRIMARY KEY (id)
			);
		`,
		`
			DROP TABLE IF EXISTS roles;
		`,
	)
}

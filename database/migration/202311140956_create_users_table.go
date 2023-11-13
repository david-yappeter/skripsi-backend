package migration

func init() {
	sourceDriver.append(
		202311140956,
		`
			CREATE TABLE IF NOT EXISTS users (
				id char(36) NOT NULL,
				username varchar(255) NOT NULL,
				name varchar(255) NOT NULL,
				password varchar(255) NOT NULL,
				is_active bool NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT users_pk PRIMARY KEY (id),
				CONSTRAINT users_uk UNIQUE (username)
			);
		`,
		`
			DROP TABLE IF EXISTS users;
		`,
	)
}

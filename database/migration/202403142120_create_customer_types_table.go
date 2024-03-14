package migration

func init() {
	sourceDriver.append(
		202403142120,
		`
			CREATE TABLE IF NOT EXISTS customer_types (
				id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT customer_types_pk PRIMARY KEY (id),
				CONSTRAINT customer_types_uk_name UNIQUE (name)
			);
		`,
		`
			DROP TABLE IF EXISTS customer_types;
		`,
	)
}

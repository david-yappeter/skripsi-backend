package migration

func init() {
	sourceDriver.append(
		202401051210,
		`
			CREATE TABLE IF NOT EXISTS supplier_types (
				id char(36) NOT NULL,
				name varchar (255) NOT NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT supplier_types_pk PRIMARY KEY (id),
				CONSTRAINT supplier_types_uk_name UNIQUE (name)
			);
		`,
		`
			DROP TABLE IF EXISTS supplier_types;
		`,
	)
}

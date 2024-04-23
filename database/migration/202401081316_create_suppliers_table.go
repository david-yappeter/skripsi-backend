package migration

func init() {
	sourceDriver.append(
		202401081316,
		`
			CREATE TABLE IF NOT EXISTS suppliers (
				id char(36) NOT NULL,
				supplier_type_id char(36) NOT NULL,
				code varchar(255) NOT NULL,
				name varchar(255) NOT NULL,
				is_active bool NOT NULL,
				address text NOT NULL,
				phone varchar(20) NOT NULL,
				email varchar(255) NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT suppliers_pk PRIMARY KEY (id),
				CONSTRAINT suppliers_supplier_types_fk FOREIGN KEY (supplier_type_id) REFERENCES supplier_types (id),
				CONSTRAINT suppliers_uk_code UNIQUE (code)
			);
		`,
		`
			DROP TABLE IF EXISTS suppliers;
		`,
	)
}

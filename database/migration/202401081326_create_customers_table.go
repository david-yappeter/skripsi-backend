package migration

func init() {
	sourceDriver.append(
		202401081326,
		`
			CREATE TABLE IF NOT EXISTS customers (
				id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				email varchar(255) NOT NULL,
				address text NOT NULL,
				phone varchar(20) NOT NULL,
				latitude DOUBLE PRECISION NOT NULL,
				longitude DOUBLE PRECISION NOT NULL,
				is_active bool NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT customers_pk PRIMARY KEY (id)
			);
		`,
		`
			DROP TABLE IF EXISTS customers;
		`,
	)
}

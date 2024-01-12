package migration

func init() {
	sourceDriver.append(
		202401120950,
		`
			CREATE TABLE IF NOT EXISTS balances (
				id char(36) NOT NULL,
				account_number varchar(255) NOT NULL,
				account_name varchar(255) NOT NULL,
				bank_name varchar(255) NOT NULL,
				name varchar(255) NOT NULL,
				amount decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT balances_pk PRIMARY KEY (id)
			);
		`,
		`
			DROP TABLE IF EXISTS balances;
		`,
	)
}

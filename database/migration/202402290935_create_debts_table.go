package migration

func init() {
	sourceDriver.append(
		202402290935,
		`
			CREATE TABLE IF NOT EXISTS debts (
				id char(36) NOT NULL,
				debt_source varchar(255) NOT NULL,
				debt_source_identifier varchar(255) NOT NULL,
				due_date date NULL,
				status varchar(255) NOT NULL,
				amount decimal(16,2) NOT NULL,
				remaining_amount decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT debts_pk PRIMARY KEY (id)
			);
		`,
		`
			DROP TABLE IF EXISTS debts;
		`,
	)
}

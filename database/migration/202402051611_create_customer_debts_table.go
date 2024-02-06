package migration

func init() {
	sourceDriver.append(
		202402051611,
		`
			CREATE TABLE IF NOT EXISTS customer_debts (
				id char(36) NOT NULL,
				customer_id char(36) NOT NULL,
				debt_source varchar(255) NOT NULL,
				debt_source_id varchar(255) NOT NULL,
				due_date date NULL,
				status varchar(255) NOT NULL,
				amount decimal(16,2) NOT NULL,
				remaining_amount decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT customer_debts_pk PRIMARY KEY (id),
				CONSTRAINT customer_debts_customers_fk FOREIGN KEY (customer_id) REFERENCES customers (id)
			); 
		`,
		`
			DROP TABLE IF EXISTS customer_debts;
		`,
	)
}

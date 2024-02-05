package migration

func init() {
	sourceDriver.append(
		202402051614,
		`
			CREATE TABLE IF NOT EXISTS customer_payments (
				id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				customer_debt_id char(36) NOT NULL,
				amount decimal(16,2) NOT NULL,
				description text NULL,
				paid_at timestamp NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT customer_payments_pk PRIMARY KEY (id),
				CONSTRAINT customer_payments_customer_debts_fk FOREIGN KEY (customer_debt_id) REFERENCES customer_debts (id)
			);
		`,
		`
			DROP TABLE IF EXISTS customer_payments;
		`,
	)
}

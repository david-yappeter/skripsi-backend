package migration

func init() {
	sourceDriver.append(
		202402230903,
		`
			CREATE TABLE IF NOT EXISTS transactions (
				id char(36) NOT NULL,
				cashier_session_id char(36) NOT NULL,
				status varchar(255) NOT NULL,
				total decimal(16,2) NOT NULL, 
				payment_at timestamp NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT transactions_pk PRIMARY KEY (id),
				CONSTRAINT transactions_fk FOREIGN KEY (cashier_session_id) REFERENCES cashier_sessions (id)
			);
		`,
		`
			DROP TABLE IF EXISTS transactions;
		`,
	)
}

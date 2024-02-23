package migration

func init() {
	sourceDriver.append(
		202402231436,
		`
			CREATE TABLE IF NOT EXISTS transaction_payments (
				id char(36) NOT NULL,
				transaction_id char(36) NOT NULL,
				status varchar(255) NOT NULL,
				payment_type varchar(255) NOT NULL,
				reference_number varchar(255) NULL,
				total decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT transaction_payments_pk PRIMARY KEY (id),
				CONSTRAINT transaction_payments_transactions_fk FOREIGN KEY (transaction_id) REFERENCES transactions (id),
				CONSTRAINT transaction_payments_uk_transaction_id UNIQUE (transaction_id)
			);
		`,
		`
			DROP TABLE IF EXISTS transaction_payments;
		`,
	)
}

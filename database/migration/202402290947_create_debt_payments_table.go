package migration

func init() {
	sourceDriver.append(
		202402290947,
		`
			CREATE TABLE IF NOT EXISTS debt_payments (
				id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				image_file_id char(36) NOT NULL,
				debt_id char(36) NOT NULL,
				amount decimal(16,2) NOT NULL,
				description text NULL,
				paid_at timestamp NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT debt_payments_pk PRIMARY KEY (id),
				CONSTRAINT debt_payments_users_fk FOREIGN KEY (user_id) REFERENCES users (id),
				CONSTRAINT debt_payments_files_fk FOREIGN KEY (image_file_id) REFERENCES files (id),
				CONSTRAINT debt_payments_debts_fk FOREIGN KEY (debt_id) REFERENCES debts (id)
			);
		`,
		`
			DROP TABLE IF EXISTS debt_payments;
		`,
	)
}

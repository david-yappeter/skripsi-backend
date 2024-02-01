package migration

func init() {
	sourceDriver.append(
		202401221030,
		`
			CREATE TABLE IF NOT EXISTS cashier_sessions (
				id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				status varchar(255) NOT NULL,
				starting_cash decimal(16,2) NOT NULL,
				ending_cash decimal(16,2) NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT cashier_sessions_pk PRIMARY KEY (id),
				CONSTRAINT cashier_sessions_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS cashier_sessions;
		`,
	)
}

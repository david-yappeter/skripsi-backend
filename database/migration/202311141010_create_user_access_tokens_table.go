package migration

func init() {
	sourceDriver.append(
		202311141010,
		`
			CREATE TABLE IF NOT EXISTS user_access_tokens (
				id char(100) NOT NULL,
				user_id char(36) NOT NULL,
				revoked boolean NOT NULL,
				created_at timestamp NULL,
				updated_at timestamp NULL,
				expired_at timestamp NOT NULL,
				CONSTRAINT user_access_tokens_pk PRIMARY KEY (id),
				CONSTRAINT user_access_tokens_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS user_access_tokens;
		`,
	)
}

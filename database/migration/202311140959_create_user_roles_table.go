package migration

func init() {
	sourceDriver.append(
		202311140959,
		`
			CREATE TABLE IF NOT EXISTS user_roles (
				user_id char(36) NOT NULL,
				role_id char(36) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT user_roles_users_fk FOREIGN KEY (user_id) REFERENCES users (id),
				CONSTRAINT user_roles_roles_fk FOREIGN KEY (role_id) REFERENCES roles (id),
				CONSTRAINT user_roles_pk PRIMARY KEY (user_id,role_id)
			);
		`,
		`
			DROP TABLE IF EXISTS user_roles;
		`,
	)
}

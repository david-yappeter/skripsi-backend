package migration

func init() {
	sourceDriver.append(
		202311131542,
		`
			CREATE TABLE IF NOT EXISTS role_permissions (
				role_id char(36) NOT NULL,
				permission_id char(36) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT role_permission_pk PRIMARY KEY (role_id,permission_id),
				CONSTRAINT role_permission_permission_fk FOREIGN KEY (permission_id) REFERENCES permissions (id),
				CONSTRAINT role_permission_roles_fk FOREIGN KEY (role_id) REFERENCES roles (id)
			);
		`,
		`
			DROP TABLE IF EXISTS role_permissions;
		`,
	)
}

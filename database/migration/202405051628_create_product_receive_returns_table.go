package migration

func init() {
	sourceDriver.append(
		202405051628,
		`
			CREATE TABLE IF NOT EXISTS product_receive_returns (
				id char(36) NOT NULL,
				product_receive_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_receive_returns_pk PRIMARY KEY (id),
				CONSTRAINT product_receive_returns_product_receives_fk FOREIGN KEY (product_receive_id) REFERENCES product_receives (id),
				CONSTRAINT product_receive_returns_users_fk FOREIGN KEY (user_id) REFERENCES users (id),
				CONSTRAINT product_receive_returns_product_receive_id_uk UNIQUE (product_receive_id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_receive_returns;
		`,
	)
}

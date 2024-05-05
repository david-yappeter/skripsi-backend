package migration

func init() {
	sourceDriver.append(
		202405051623,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_returns (
				id char(36) NOT NULL,
				delivery_order_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_returns_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_returns_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES delivery_orders (id),
				CONSTRAINT delivery_order_returns_users_fk FOREIGN KEY (user_id) REFERENCES users (id),
				CONSTRAINT delivery_order_returns_delivery_order_id_uk UNIQUE (delivery_order_id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_returns;
		`,
	)
}

package migration

func init() {
	sourceDriver.append(
		202403201525,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_positions (
				id char(36) NOT NULL,
				delivery_order_id char(36) NOT NULL,
				driver_user_id char(36) NOT NULL,
				latitude DOUBLE PRECISION NOT NULL,
				longitude DOUBLE PRECISION NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_positions_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_positions_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES delivery_orders (id),
				CONSTRAINT delivery_order_positions_users_fk FOREIGN KEY (driver_user_id) REFERENCES users (id),
				CONSTRAINT delivery_order_positions_uk_delivery_order_id UNIQUE (delivery_order_id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_positions;
		`,
	)
}

package migration

func init() {
	sourceDriver.append(
		202401191648,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_drivers (
				id char(36) NOT NULL,
				delivery_order_id char(36) NOT NULL,
				driver_user_id char(36) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_drivers_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_drivers_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES delivery_orders (id),
				CONSTRAINT delivery_order_drivers_users_fk FOREIGN KEY (driver_user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_drivers
		`,
	)
}

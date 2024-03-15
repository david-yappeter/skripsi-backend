package migration

func init() {
	sourceDriver.append(
		202401191647,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_items (
				id char(36) NOT NULL,
				delivery_order_id char(36) NOT NULL,
				product_unit_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				qty decimal(16,2) NOT NULL,
				price_per_unit decimal(16,2) NOT NULL,
				discount_per_unit decimal(16,2) NOT NULL,m
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_items_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_items_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES delivery_orders (id),
				CONSTRAINT delivery_order_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id),
				CONSTRAINT delivery_order_items_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_items;
		`,
	)
}

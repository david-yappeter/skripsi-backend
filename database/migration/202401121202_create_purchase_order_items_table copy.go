package migration

func init() {
	sourceDriver.append(
		202401121202,
		`
			CREATE TABLE IF NOT EXISTS purchase_order_items (
				id char(36) NOT NULL,
				purchase_order_id char(36) NOT NULL,
				product_unit_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				qty decimal(16,2) NOT NULL,
				scale_to_base decimal(16,2) NOT NULL,
				price_per_unit decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT purchase_order_items_pk PRIMARY KEY (id),
				CONSTRAINT purchase_order_items_purchase_orders_fk FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders (id),
				CONSTRAINT purchase_order_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id),
				CONSTRAINT purchase_order_items_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS purchase_orders;
		`,
	)
}

package migration

func init() {
	sourceDriver.append(
		202401160910,
		`
			CREATE TABLE IF NOT EXISTS product_receive_items (
				id char(36) NOT NULL,
				product_receive_id char(36) NOT NULL,
				product_unit_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				qty_eligible decimal(16,2) NOT NULL,
				qty_received decimal(16,2) NOT NULL,
				scale_to_base decimal(16,2) NOT NULL,
				price_per_unit decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_receive_items_pk PRIMARY KEY (id),
				CONSTRAINT product_receive_items_product_receives_fk FOREIGN KEY (product_receive_id) REFERENCES product_receives (id),
				CONSTRAINT product_receive_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id),
				CONSTRAINT product_receive_items_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_receive_items;
		`,
	)
}

package migration

func init() {
	sourceDriver.append(
		202406061501,
		`
			CREATE TABLE IF NOT EXISTS product_return_items (
				id char(36) NOT NULL,
				product_return_id char(36) NOT NULL,
				product_unit_id char(36) NOT NULL,
				qty decimal(16,2) NOT NULL,
				scale_to_base decimal(16,2) NOT NULL,
				base_cost_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_return_items_pk PRIMARY KEY (id),
				CONSTRAINT product_return_items_product_returns_fk FOREIGN KEY (product_return_id) REFERENCES product_returns (id),
				CONSTRAINT product_return_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_return_items;
		`,
	)
}

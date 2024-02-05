package migration

func init() {
	sourceDriver.append(
		202401221104,
		`
			CREATE TABLE IF NOT EXISTS cart_items (
				id char(36) NOT NULL,
				cart_id char(36) NOT NULL, 
				product_unit_id char(36) NOT NULL, 
				qty decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT cart_items_pk PRIMARY KEY (id),
				CONSTRAINT cart_items_carts_fk FOREIGN KEY (cart_id) REFERENCES carts (id),
				CONSTRAINT cart_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id),
				CONSTRAINT cart_items_uk_1 UNIQUE (cart_id, product_unit_id)
			);
		`,
		`
			DROP TABLE IF EXISTS cart_items;
		`,
	)
}

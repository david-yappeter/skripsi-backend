package migration

func init() {
	sourceDriver.append(
		202403151330,
		`
			CREATE TABLE IF NOT EXISTS product_discounts (
				id char(36) NOT NULL,
				product_id char(36) NOT NULL,
				minimum_qty decimal(16,2) NOT NULL,
				is_active bool NOT NULL,
				discount_percentage decimal(16,2) NULL,
				discount_amount decimal(16,2) NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_discounts_pk PRIMARY KEY (id),
				CONSTRAINT product_discounts_products_fk FOREIGN KEY (product_id) REFERENCES products (id),
				CONSTRAINT product_discounts_uk_product_id UNIQUE (product_id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_discounts;
		`,
	)
}

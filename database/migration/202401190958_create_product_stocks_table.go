package migration

func init() {
	sourceDriver.append(
		202401190958,
		`
			CREATE TABLE IF NOT EXISTS product_stocks (
				id char(36) NOT NULL,
				product_id char(36) NOT NULL,
				qty decimal(16,2) NOT NULL,
				base_cost_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_stocks_pk PRIMARY KEY (id),
				CONSTRAINT product_stocks_products_fk FOREIGN KEY (product_id) REFERENCES products (id),
				CONSTRAINT product_stocks_uk_1 UNIQUE (product_id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_stocks;
		`,
	)
}

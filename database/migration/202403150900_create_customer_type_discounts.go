package migration

func init() {
	sourceDriver.append(
		202403150900,
		`
			CREATE TABLE IF NOT EXISTS customer_type_discounts (
				id char(36) NOT NULL,
				product_id char(36) NOT NULL,
				customer_type_id char(36) NOT NULL,
				is_active bool NOT NULL,
				discount_percentage decimal(16,2) NULL,
				discount_amount decimal(16,2) NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT customer_type_discounts_pk PRIMARY KEY (id),
				CONSTRAINT customer_type_discounts_customer_types_fk FOREIGN KEY (customer_type_id) REFERENCES customer_types (id),
				CONSTRAINT customer_type_discounts_products_fk FOREIGN KEY (product_id) REFERENCES products (id),
				CONSTRAINT customer_type_discounts_uk_1 UNIQUE (product_id, customer_type_id)
			);
		`,
		`
			DROP TABLE IF EXISTS customer_type_discounts;
		`,
	)
}

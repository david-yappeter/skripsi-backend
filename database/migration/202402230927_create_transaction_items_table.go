package migration

func init() {
	sourceDriver.append(
		202402230927,
		`
			CREATE TABLE IF NOT EXISTS transaction_items (
				id char(36) NOT NULL,
				transaction_id char(36) NOT NULL,
				product_unit_id char(36) NOT NULL,
				qty decimal(16,2) NOT NULL,
				price_per_unit decimal(16,2) NOT NULL,
				discount_per_unit decimal(16,2) NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT transaction_items_pk PRIMARY KEY (id),
				CONSTRAINT transaction_items_transactions_fk FOREIGN KEY (transaction_id) REFERENCES transactions (id),
				CONSTRAINT transaction_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id)
			);
		`,
		`
			DROP TABLE IF EXISTS transaction_items;
		`,
	)
}

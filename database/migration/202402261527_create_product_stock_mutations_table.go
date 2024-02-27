package migration

func init() {
	sourceDriver.append(
		202402261527,
		`
			CREATE TABLE IF NOT EXISTS product_stock_mutations (
				id char(36) NOT NULL,
				product_unit_id char(36) NOT NULL,
				type varchar(255) NOT NULL,
				identifier_id varchar(255) NOT NULL,
				qty decimal(16,2) NOT NULL,
				scale_to_base decimal(16,2) NOT NULL,
				base_qty_left decimal(16,2) NOT NULL,
				base_cost_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_stock_mutations_pk PRIMARY KEY (id),
				CONSTRAINT product_stock_mutations_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_stock_mutations;
		`,
	)
}

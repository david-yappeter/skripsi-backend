package migration

func init() {
	sourceDriver.append(
		202405212245,
		`
			CREATE TABLE IF NOT EXISTS product_stock_adjustments (
				id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				product_stock_id char(36) NOT NULL,
				previous_qty decimal(16,2) NOT NULL,
				updated_qty decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_stock_adjustments_pk PRIMARY KEY (id),
				CONSTRAINT product_stock_adjustments_users_fk FOREIGN KEY (user_id) REFERENCES users (id),
				CONSTRAINT product_stock_adjustments_product_stocks_fk FOREIGN KEY (product_stock_id) REFERENCES product_stocks (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_stock_adjustments;
		`,
	)
}

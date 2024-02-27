package migration

func init() {
	sourceDriver.append(
		202402271401,
		`
			CREATE TABLE IF NOT EXISTS transaction_item_costs (
				id char(36),
				transaction_item_id char(36) NOT NULL,
				qty decimal(16,2) NOT NULL,
				base_cost_price decimal(16,2) NOT NULL,
				total_cost_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT transaction_item_costs_pk PRIMARY KEY (id),
				CONSTRAINT transaction_item_costs_transaction_items_fk FOREIGN KEY (transaction_item_id) REFERENCES transaction_items (id)
			);
		`,
		`
			DROP TABLE IF EXISTS transaction_item_costs;
		`,
	)
}

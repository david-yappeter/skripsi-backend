package migration

func init() {
	sourceDriver.append(
		202402271658,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_item_costs (
				id char(36) NOT NULL,
				delivery_order_item_id char(36) NOT NULL,
				qty decimal(16,2) NOT NULL,
				base_cost_price decimal(16,2) NOT NULL,
				total_cost_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_item_costs_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_item_costs_delivery_order_items_fk FOREIGN KEY (delivery_order_item_id) REFERENCES delivery_order_items (id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_item_costs;
		`,
	)
}

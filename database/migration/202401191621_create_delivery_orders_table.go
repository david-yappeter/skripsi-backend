package migration

func init() {
	sourceDriver.append(
		202401191621,
		`
			CREATE TABLE IF NOT EXISTS delivery_orders (
				id char(36) NOT NULL,
				customer_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				invoice_number varchar(255) NOT NULL,
				date date NOT NULL,
				status varchar(255) NOT NULL,
				total_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_orders_pk PRIMARY KEY (id),
				CONSTRAINT delivery_orders_customers_fk FOREIGN KEY (customer_id) REFERENCES customers (id),
				CONSTRAINT delivery_orders_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_orders;
		`,
	)
}

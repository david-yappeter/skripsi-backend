package migration

func init() {
	sourceDriver.append(
		202401121201,
		`
			CREATE TABLE IF NOT EXISTS purchase_orders (
				id char(36) NOT NULL,
				supplier_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				invoice_number varchar(255) NOT NULL,
				date date NOT NULL,
				status varchar(255) NOT NULL,
				total_estimated_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT purchase_orders_pk PRIMARY KEY (id),
				CONSTRAINT purchase_orders_suppliers_fk FOREIGN KEY (supplier_id) REFERENCES suppliers (id),
				CONSTRAINT purchase_orders_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS purchase_orders;
		`,
	)
}

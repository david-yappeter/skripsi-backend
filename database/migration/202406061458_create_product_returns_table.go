package migration

func init() {
	sourceDriver.append(
		202406061458,
		`
			CREATE TABLE IF NOT EXISTS product_returns (
				id char(36) NOT NULL,
				supplier_id char(36) NOT NULL,
				user_id char(36) NOT NULL,
				invoice_number varchar(255) NOT NULL,
				date date NOT NULL,
				status varchar(255) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_returns_pk PRIMARY KEY (id),
				CONSTRAINT product_returns_suppliers_fk FOREIGN KEY (supplier_id) REFERENCES suppliers (id),
				CONSTRAINT product_returns_users_fk FOREIGN KEY (user_id) REFERENCES users (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_returns;
		`,
	)
}

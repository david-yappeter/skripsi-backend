package migration

func init() {
	sourceDriver.append(
		202403142122,
		`
			ALTER TABLE customers
			ADD COLUMN customer_type_id char(36) NULL,
			ADD CONSTRAINT customers_customer_types_fk FOREIGN KEY (customer_type_id) REFERENCES customer_types (id);
		`,
		`
			ALTER TABLE customers
			DROP CONSTRAINT customers_customer_types_fk,
			DROP COLUMN customer_type_id;
		`,
	)
}

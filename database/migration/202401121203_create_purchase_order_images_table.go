package migration

func init() {
	sourceDriver.append(
		202401121203,
		`
		CREATE TABLE IF NOT EXISTS purchase_order_images (
			id char(36) NOT NULL,
			purchase_order_id char(36) NOT NULL,
			file_id char(36) NOT NULL,
			description text NULL,
			created_at timestamp NOT NULL,
			updated_at timestamp NOT NULL,
			CONSTRAINT purchase_order_images_pk PRIMARY KEY (id),
			CONSTRAINT purchase_order_images_purchase_orders_fk FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders (id),
			CONSTRAINT purchase_order_images_files_fk FOREIGN KEY (file_id) REFERENCES files (id)
		);
		`,
		`
			DROP TABLE IF EXISTS purchase_order_images;
		`,
	)
}

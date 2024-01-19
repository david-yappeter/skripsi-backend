package migration

func init() {
	sourceDriver.append(
		202401191648,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_images (
				id char(36) NOT NULL,
				delivery_order_id char(36) NOT NULL,
				file_id char(36) NOT NULL,
				description varchar(255) NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_images_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_images_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES delivery_orders (id),
				CONSTRAINT delivery_order_images_files_fk FOREIGN KEY (file_id) REFERENCES files (id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_images;
		`,
	)
}

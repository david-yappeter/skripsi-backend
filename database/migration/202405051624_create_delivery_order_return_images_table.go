package migration

func init() {
	sourceDriver.append(
		202405051624,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_return_images (
				id char(36) NOT NULL,
				delivery_order_return_id char(36) NOT NULL,
				file_id char(36) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_return_images_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_return_images_delivery_order_returns_fk FOREIGN KEY (delivery_order_return_id) REFERENCES delivery_order_returns (id),
				CONSTRAINT delivery_order_return_images_files_fk FOREIGN KEY (file_id) REFERENCES files (id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_return_images;
		`,
	)
}

package migration

func init() {
	sourceDriver.append(
		202401161435,
		`
			CREATE TABLE IF NOT EXISTS product_receive_images (
				id char(36) NOT NULL,
				product_receive_id char(36) NOT NULL,
				file_id char(36) NOT NULL,
				description text NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_receives_images_pk PRIMARY KEY (id),
				CONSTRAINT product_receives_images_product_receives_fk FOREIGN KEY (product_receive_id) REFERENCES product_receives (id),
				CONSTRAINT product_receives_images_files_fk FOREIGN KEY (file_id) REFERENCES files (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_receive_images;
		`,
	)
}

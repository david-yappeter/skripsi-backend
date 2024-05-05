package migration

func init() {
	sourceDriver.append(
		202405051629,
		`
			CREATE TABLE IF NOT EXISTS product_receive_return_images (
				id char(36) NOT NULL,
				product_receive_return_id char(36) NOT NULL,
				file_id char(36) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_receive_return_images_pk PRIMARY KEY (id),
				CONSTRAINT product_receive_return_images_files_fk FOREIGN KEY (file_id) REFERENCES files (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_receive_return_images;
		`,
	)
}

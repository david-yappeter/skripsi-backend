package migration

func init() {
	sourceDriver.append(
		202406061503,
		`
			CREATE TABLE IF NOT EXISTS product_return_images (
				id char(36) NOT NULL,
				product_return_id char(36) NOT NULL,
				file_id char(36) NOT NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_return_images_pk PRIMARY KEY (id),
				CONSTRAINT product_return_images_product_returns_fk FOREIGN KEY (product_return_id) REFERENCES product_returns (id),
				CONSTRAINT product_return_images_files_fk FOREIGN KEY (file_id) REFERENCES files (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_return_images;
		`,
	)
}

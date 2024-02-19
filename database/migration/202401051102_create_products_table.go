package migration

func init() {
	sourceDriver.append(
		202401051102,
		`
			CREATE TABLE IF NOT EXISTS products (
				id char(36) NOT NULL,
				image_file_id char(36) NOT NULL,
				name varchar(255) NOT NULL,
				description text NULL,
				price decimal(16,2) NULL,
				is_active bool NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT products_pk PRIMARY KEY (id),
				CONSTRAINT products_files_fk FOREIGN KEY (image_file_id) REFERENCES files (id),
				CONSTRAINT products_uk_name UNIQUE (name)
			);
		`,
		`
			DROP TABLE IF EXISTS products;
		`,
	)
}

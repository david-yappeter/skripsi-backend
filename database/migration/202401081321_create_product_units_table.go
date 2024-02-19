package migration

func init() {
	sourceDriver.append(
		202401081321,
		`
			CREATE TABLE IF NOT EXISTS product_units (
				id char(36) NOT NULL,
				to_unit_id char(36) NULL,
				image_file_id char(36) NULL,
				unit_id char(36) NOT NULL,
				product_id char(36) NOT NULL,
				scale decimal(16,2) NOT NULL,
				scale_to_base decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT product_units_pk PRIMARY KEY (id),
				CONSTRAINT product_units_files_fk FOREIGN KEY (image_file_id) REFERENCES files (id),
				CONSTRAINT product_units_units_fk_1 FOREIGN KEY (to_unit_id) REFERENCES units (id),
				CONSTRAINT product_units_units_fk_2 FOREIGN KEY (unit_id) REFERENCES units (id),
				CONSTRAINT product_units_products_fk FOREIGN KEY (product_id) REFERENCES products (id)
			);
		`,
		`
			DROP TABLE IF EXISTS product_units;
		`,
	)
}

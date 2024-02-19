package migration

func init() {
	sourceDriver.append(
		202402161456,
		`
			CREATE TABLE IF NOT EXISTS tiktok_products (
				tiktok_product_id varchar(255) NOT NULL,
				product_id char(36) NOT NULL,
				status varchar(255) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT tiktok_products_pk PRIMARY KEY (tiktok_product_id),
				CONSTRAINT tiktok_products_products_fk FOREIGN KEY (product_id) REFERENCES products (id),
				CONSTRAINT tiktok_products_product_uk UNIQUE (product_id)
			);
		`,
		`
			DROP TABLE IF EXISTS tiktok_products;
		`,
	)
}

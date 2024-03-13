package migration

func init() {
	sourceDriver.append(
		202403081533,
		`
			CREATE TABLE IF NOT EXISTS shop_order_items (
				id char(36) NOT NULL,
				shop_order_id char(36) NOT NULL,
				product_unit_id char(36) NOT NULL,
				platform_product_id varchar(255) NOT NULL,
				image_link varchar(255) NULL,
				quantity decimal(16,2) NOT NULL,
				original_price decimal(16,2) NOT NULL,
				sale_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT shop_order_items_pk PRIMARY KEY (id),
				CONSTRAINT shop_order_items_shop_orders_fk FOREIGN KEY (shop_order_id) REFERENCES shop_orders (id),
				CONSTRAINT shop_order_items_product_units_fk FOREIGN KEY (product_unit_id) REFERENCES product_units (id)
			);
		`,
		`
			DROP TABLE IF EXISTS shop_order_items;
		`,
	)
}

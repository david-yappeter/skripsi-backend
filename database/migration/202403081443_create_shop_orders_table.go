package migration

func init() {
	sourceDriver.append(
		202403081443,
		`
			CREATE TABLE IF NOT EXISTS shop_orders (
				id char(36) NOT NULL,
				tracking_number varchar(255) NULL,
				platform_type varchar(255) NOT NULL,
				tracking_status varchar(255) NOT NULL,
				recipient_name varchar(255) NOT NULL,
				recipient_full_address text NOT NULL,
				recipient_phone_number varchar(20) NOT NULL,
				original_price decimal(16,2) NOT NULL,
				sale_price decimal(16,2) NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT shop_orders_pk PRIMARY KEY (id),
				CONSTRAINT shop_orders_uk_tracking_number UNIQUE (tracking_number)
			);
		`,
		`
			DROP TABLE IF EXISTS shop_orders;
		`,
	)
}

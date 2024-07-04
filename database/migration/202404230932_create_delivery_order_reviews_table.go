package migration

func init() {
	sourceDriver.append(
		202404230932,
		`
			CREATE TABLE IF NOT EXISTS delivery_order_reviews (
				id char(36) NOT NULL,
				delivery_order_id char(36) NOT NULL,
				type varchar(255) NOT NULL,
				star_rating int NOT NULL,
				description text NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT delivery_order_reviews_pk PRIMARY KEY (id),
				CONSTRAINT delivery_order_reviews_delivery_orders_fk FOREIGN KEY (delivery_order_id) REFERENCES delivery_orders (id),
				CONSTRAINT delivery_order_delivery_order_id_uk UNIQUE (delivery_order_id)
			);
		`,
		`
			DROP TABLE IF EXISTS delivery_order_reviews;
		`,
	)
}

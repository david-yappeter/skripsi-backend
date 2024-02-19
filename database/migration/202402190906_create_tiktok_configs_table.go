package migration

func init() {
	sourceDriver.append(
		202402190906,
		`
			CREATE TABLE IF NOT EXISTS tiktok_configs (
				app_key varchar(255) NOT NULL,
				app_secret varchar(255) NOT NULL,
				warehouse_id varchar(255) NOT NULL,
				shop_id varchar(255) NOT NULL,
				shop_cipher varchar(255) NOT NULL,
				access_token varchar(255) NULL,
				refresh_token varchar(255) NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT tiktok_configs_pk PRIMARY KEY (app_key)
			);
		`,
		`
			DROP TABLE IF EXISTS tiktok_configs;
		`,
	)
}

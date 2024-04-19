package migration

func init() {
	sourceDriver.append(
		202404191453,
		`
			CREATE TABLE IF NOT EXISTS shopee_configs (
				partner_id varchar(255) NOT NULL,
				partner_key varchar(255) NOT NULL,
				access_token varchar(255) NULL,
				refresh_token varchar(255) NULL,
				CONSTRAINT shopee_configs_pk PRIMARY KEY (partner_id)
			);
		`,
		`
			DROP TABLE IF EXISTS shopee_configs;
		`,
	)
}

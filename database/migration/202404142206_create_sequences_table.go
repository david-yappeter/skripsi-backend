package migration

func init() {
	sourceDriver.append(
		202404142206,
		`
			CREATE TABLE IF NOT EXISTS sequences (
				id char(36) NOT NULL,
				unique_identifier varchar(255) NOT NULL,
				sequence int NOT NULL,
				created_at timestamp NOT NULL,
				updated_at timestamp NOT NULL,
				CONSTRAINT sequences_pk PRIMARY KEY (id),
				CONSTRAINT sequences_uk_1 UNIQUE(unique_identifier, sequence)
			);
		`,
		`
			DROP TABLE IF EXISTS sequences;
		`,
	)
}

package store

func autoMigrate() {
	db.AutoMigrate(Account{},
		Unit{})
	migrateScript()
}

func migrateScript() {
	db.Exec(`CREATE EXTENSION IF NOT EXISTS pg_trgm;
	ALTER TABLE account ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(acc_name, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(acc_group, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_account_text_search ON account USING GIN(text_search);
`)
}

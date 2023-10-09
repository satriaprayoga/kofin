package store

func autoMigrate() {
	db.AutoMigrate(Account{},
		Unit{},
		KProgram{},
		Kegiatan{})
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
	ALTER TABLE k_program ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(program_name, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(unit_name, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_program_text_search ON k_program USING GIN(text_search);
	ALTER TABLE kegiatan ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(kegiatan_name, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(program_name, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_kegiatan_text_search ON kegiatan USING GIN(text_search);
`)
}

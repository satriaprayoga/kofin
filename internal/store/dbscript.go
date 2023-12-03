package store

func autoMigrate() {
	db.AutoMigrate(Account{},
		Unit{},
		KProgram{},
		Kegiatan{},
		Budget{},
		Expend{},
		ExpendProgram{},
		ExpendKegiatan{},
		ExpendAccount{},
		ExpendObject{},
		KUser{},
		KRole{})
	migrateScript()
	seedData()
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
	ALTER TABLE expend_program ADD text_search tsvector 
		GENERATED ALWAYS AS	(
			setweight(to_tsvector('indonesian', coalesce(program_name, '')), 'A') || ' ' ||
			setweight(to_tsvector('indonesian', coalesce(slug, '')), 'B') || ' ' || 
			setweight(to_tsvector('indonesian', coalesce(unit_name, '')), 'C') :: tsvector
		) STORED;
	CREATE INDEX idx_expend_program_text_search ON expend_program USING GIN(text_search);
`)
}

func seedData() {
	var roles = []KRole{{Name: "admin", Description: "Administrator role"}, {Name: "anggaran", Description: "Anggaran role"}, {Name: "anonymous", Description: "Unauthenticated role"}}
	var user = []KUser{{Username: "ngadimin", Fullname: "Admin Suradmin", Password: "asdqwe123", RoleID: 1}}

	db.Save(&roles)

	var exists KUser
	err := db.Where("role_id=?", 1).First(&exists)
	if err.RowsAffected < 1 {
		db.Save(&user)
	}

	// var budget Budget
	// year := time.Now().Year()
	// err = db.Where("budget_year=?", year).First(&budget)
	// if err.RowsAffected < 1 {
	// 	db.Save(&Budget{
	// 		BudgetYear:   year + 1,
	// 		BudgetStatus: 0,
	// 		BudgetDate:   time.Now(),
	// 	})
	// }
}

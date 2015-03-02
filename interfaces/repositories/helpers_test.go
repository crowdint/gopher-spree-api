package repositories

func ResetDB() {
	Spree_db.Rollback()
	Spree_db.Close()
}

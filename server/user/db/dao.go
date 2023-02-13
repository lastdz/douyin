package db

func InsertUser(name string, passwd string) (int64, error) {
	user := &UserModel{Name: name, Password: passwd}
	tx := db.Create(user)
	return user.Id, tx.Error
}

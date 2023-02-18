package db

func InsertUser(name string, passwd string) (int64, error) {
	user := &UserModel{Name: name, Password: passwd}
	tx := db.Create(user)
	return user.Id, tx.Error
}

func ExistByNameAndPasswd(name string, passwd string) (*UserModel, error) {
	var user UserModel
	result := db.Where("name = ? AND password = ?", name, passwd).Take(&user)
	return &user, result.Error
}

func GetByName(name string) (*UserModel, error) {
	var user UserModel
	result := db.Where("name = ?", name).Take(&user)
	return &user, result.Error
}

func GetById(id int64) (*UserModel, error) {
	var user UserModel
	result := db.Where("id = ?", id).Take(&user)
	return &user, result.Error
}

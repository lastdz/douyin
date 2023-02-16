package db

func Favorite(fav *FavoriteDb) error {
	tx := mydb.Create(fav)
	return tx.Error
}
func UnFavorite(fav *FavoriteDb) error {
	tx := mydb.Delete(&FavoriteDb{
		token: fav.token,
	})
	return tx.Error
}
func GetAllFavorite(token string) (favs []FavoriteDb) {
	mydb.Table("favorite_dbs").Where("token=?", token).Find(&favs)
	return
}

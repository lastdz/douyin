package db

func InsertComment(comment *CommentDb) error {
	tx := mydb.Create(comment)
	return tx.Error
}

func DeleteComment(id int64) error {
	tx := mydb.Delete(&CommentDb{Id: id})
	return tx.Error
}
func GetAllComment(vedioid int) (commentdbs []CommentDb) {
	mydb.Table("comment_dbs").Where("vedioid=?", vedioid).Find(&commentdbs)
	return commentdbs
}

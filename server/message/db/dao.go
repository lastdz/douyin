package db

func InsertMessage(message *MessageDb) error {
	tx := myDb.Create(message)
	return tx.Error
}

func GetAllMessage(uid int64, toUserId int64) (messages []MessageDb, err error) {
	res := myDb.Table("message_dbs").Where("uid=?", uid).Where("to_user_id=?", toUserId).Find(&messages)
	if res.Error != nil {
		return nil, res.Error
	}
	return messages, nil
}

package db

import (
	"errors"
)

// 添加一条关系
func InsertRelation(relation *RelationDb) error {
	a := &RelationDb{}
	mydb.Where("userid = ? and followerid = ?", relation.Userid, relation.Followerid).First(a)
	//fmt.Println(first.Error)
	if a.Userid == relation.Userid && a.Followerid == relation.Followerid {
		return errors.New("exist")
	}
	tx := mydb.Create(relation)
	return tx.Error
}

// 删除一条关系
func DeleteRelation(relation *RelationDb) error {
	tx := mydb.Where("userid = ? and followerid = ?", relation.Userid, relation.Followerid).Delete(&RelationDb{})
	return tx.Error
}

// 获得id为Userid的人的关注id列表
func GetFollowListId(Userid int) ([]int64, error) {
	var Userids []int64
	tx := mydb.Table("relation_dbs").Select("userid").Where("followerid = ?", Userid).Find(&Userids)
	return Userids, tx.Error
}

// 获得id为Userid的人的粉丝id列表
func GetFollowerListId(Userid int) ([]int64, error) {
	var Userids []int64
	tx := mydb.Table("relation_dbs").Select("followerid").Where("userid = ?", Userid).Find(&Userids)
	return Userids, tx.Error
}

// 获得id为Userid的人的关注数量
func GetFollowCount(Userid int64) (int64, error) {
	var count int64
	tx := mydb.Table("relation_dbs").Select("*").Where("followerid = ?", Userid).Count(&count)
	return count, tx.Error
}

// 获得id为Userid的人的粉丝数量
func GetFollowerCount(Userid int64) (int64, error) {
	var count int64
	tx := mydb.Table("relation_dbs").Select("*").Where("userid = ?", Userid).Count(&count)
	return count, tx.Error
}

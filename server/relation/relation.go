package relation

import (
	"errors"
	"github.com/RaymondCode/simple-demo/server/relation/db"
	"github.com/RaymondCode/simple-demo/wire"
	"strconv"
)

func RelationAction(actionargs *wire.RelationActionArgs, actionreply *wire.RelationActionReply) error {
	a := &db.RelationDb{}
	var err error
	a.Followerid, err = strconv.ParseInt(actionargs.Token, 10, 64) //功能还未实现根据token查到userid
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	a.Userid, err = strconv.ParseInt(actionargs.To_user_id, 10, 64)
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	action_type, err2 := strconv.Atoi(actionargs.Action_type)
	if err2 != nil {
		actionreply.StatusCode = -1
		return err2
	}
	if action_type != 1 && action_type != 2 {
		actionreply.StatusCode = -1
		return errors.New("type error")
	}

	if action_type == 1 {
		err = db.InsertRelation(a)
		if err != nil {
			actionreply.StatusCode = -1
			return err
		} else {
			actionreply.StatusCode = 0
		}
	} else {
		err = db.DeleteRelation(a)
		if err != nil {
			actionreply.StatusCode = -1
			return err
		} else {
			actionreply.StatusCode = 0
		}
	}
	return nil
}

func RelationFollowList(actionargs *wire.RelationListArgs, actionreply *wire.RelationListReply) error {
	var err error
	var userid int                                  //需要获取userid的关注的人列表
	var nowuserid int                               //当前登录的人的id
	nowuserid, err = strconv.Atoi(actionargs.Token) //功能还未实现:根据token查到userid
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	userid, err = strconv.Atoi(actionargs.User_id) //功能还未实现:根据token查到userid
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	UseridList, err := db.GetFollowListId(userid)       //需要获取userid的关注的人列表
	NowUseridList, err := db.GetFollowListId(nowuserid) //当前登录的人的关注的人列表
	isfollow := make(map[int64]bool, len(NowUseridList))
	for _, val := range NowUseridList {
		isfollow[val] = true
	}
	for _, val := range UseridList {
		a := wire.User{}
		a.Id = val
		if _, exist := isfollow[a.Id]; exist == true {
			a.IsFollow = true
		} else {
			a.IsFollow = false
		}
		a.FollowCount, _ = db.GetFollowCount(a.Id)
		a.FollowerCount, _ = db.GetFollowerCount(a.Id)
		//a.Name = //功能还未实现
		actionreply.User_list = append(actionreply.User_list, a)
	}
	actionreply.StatusCode = 0
	return nil
}

func RelationFollowerList(actionargs *wire.RelationListArgs, actionreply *wire.RelationListReply) error {
	var err error
	var userid int                                  //需要获取userid的关注的人列表
	var nowuserid int                               //当前登录的人的id
	nowuserid, err = strconv.Atoi(actionargs.Token) //功能还未实现:根据token查到userid
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	userid, err = strconv.Atoi(actionargs.User_id) //功能还未实现:根据token查到userid
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	UseridList, err := db.GetFollowerListId(userid)       //需要获取userid的粉丝列表
	UseridFollowList, err := db.GetFollowerListId(userid) //当前Userid的关注的人列表
	NowUseridList, err := db.GetFollowListId(nowuserid)   //当前登录的人的关注的人列表
	isfollow := make(map[int64]bool, len(NowUseridList))
	for _, val := range NowUseridList {
		isfollow[val] = true
	}
	isUserfollow := make(map[int64]bool, len(UseridFollowList))
	for _, val := range NowUseridList {
		isUserfollow[val] = true
	}
	for _, val := range UseridList {
		if _, exist := isUserfollow[val]; exist == false {
			continue
		}
		a := wire.User{}
		a.Id = val
		if _, exist := isfollow[a.Id]; exist == true {
			a.IsFollow = true
		} else {
			a.IsFollow = false
		}
		a.FollowCount, _ = db.GetFollowCount(a.Id)
		a.FollowerCount, _ = db.GetFollowerCount(a.Id)
		//a.Name = //功能还未实现
		actionreply.User_list = append(actionreply.User_list, a)
	}
	actionreply.StatusCode = 0
	return nil
}

func RelationFriendList(actionargs *wire.RelationListArgs, actionreply *wire.RelationListReply) error {
	var err error
	var userid int                                  //需要获取userid的关注的人列表
	var nowuserid int                               //当前登录的人的id
	nowuserid, err = strconv.Atoi(actionargs.Token) //功能还未实现:根据token查到userid
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	userid, err = strconv.Atoi(actionargs.User_id) //功能还未实现:根据token查到userid
	if err != nil {
		actionreply.StatusCode = -1
		return err
	}
	UseridList, err := db.GetFollowerListId(userid)     //需要获取userid的粉丝列表
	NowUseridList, err := db.GetFollowListId(nowuserid) //当前登录的人的关注的人列表
	isfollow := make(map[int64]bool, len(NowUseridList))
	for _, val := range NowUseridList {
		isfollow[val] = true
	}
	for _, val := range UseridList {
		a := wire.User{}
		a.Id = val
		if _, exist := isfollow[a.Id]; exist == true {
			a.IsFollow = true
		} else {
			a.IsFollow = false
		}
		a.FollowCount, _ = db.GetFollowCount(a.Id)
		a.FollowerCount, _ = db.GetFollowerCount(a.Id)
		//a.Name = //功能还未实现
		actionreply.User_list = append(actionreply.User_list, a)
	}
	actionreply.StatusCode = 0
	return nil
}

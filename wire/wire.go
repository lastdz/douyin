package wire

type CommentListResponse struct {
	Response
	CommentList []Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment Comment `json:"comment,omitempty"`
}

type RegisterArgs struct {
	Username string
	Password string
}

type RegisterReply struct {
	Response
	UserId int64
	Token  string
}

type LoginArgs struct {
	Username string
	Password string
}

type LoginReply struct {
	Response
	UserId int64
	Token  string
}

type GetUserArgs struct {
	UserId int64
	Token  string
}

type GetUserReply struct {
	Response
	User User `json:"user,omitempty"`
}

type ActionArgs struct {
	Token        string
	Video_id     string
	Action_Type  string
	Comment_text string
	Comment_id   string
}
type ActionReply struct {
	Status_code    int
	Status_message string
	Comment        Comment
}
type ListArgs struct {
	Token    string
	Video_id string
}
type ListReply struct {
	Status_code    int
	Status_message string
	Comment_List   []Comment
}
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty"`
	User       User   `json:"user"`
	Content    string `json:"content,omitempty"`
	CreateDate string `json:"create_date,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type Message struct {
	Id         int64  `json:"id,omitempty"`
	Content    string `json:"content,omitempty"`
	CreateTime string `json:"create_time,omitempty"`
}

type MessageSendEvent struct {
	UserId     int64  `json:"user_id,omitempty"`
	ToUserId   int64  `json:"to_user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

type MessagePushEvent struct {
	FromUserId int64  `json:"user_id,omitempty"`
	MsgContent string `json:"msg_content,omitempty"`
}

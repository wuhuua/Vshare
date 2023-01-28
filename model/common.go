package model

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

type Video struct {
	Id            int64  `json:"id,omitempty" gorm:"primarykey"`
	Author        User   `json:"author" gorm:"-"`
	UserId        int64  `json:"-" gorm:"column:userid"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty" gorm:"column:playurl"`
	CoverUrl      string `json:"cover_url,omitempty" gorm:"column:coverurl"`
	FavoriteCount int64  `json:"favorite_count,omitempty" gorm:"column:favoritecount"`
	CommentCount  int64  `json:"comment_count,omitempty" gorm:"column:commentcount"`
	IsFavorite    bool   `json:"is_favorite,omitempty" gorm:"-"`
	VideoName     string `json:"title" gorm:"column:videoname"`
}

type Comment struct {
	Id         int64  `json:"id,omitempty" gorm:"primarykey"`
	User       User   `json:"user" gorm:"-"`
	UserId     int64  `json:"-" gorm:"column:userid"`
	VideoId    int64  `json:"-" gorm:"column:videoid"`
	Content    string `json:"content,omitempty" gorm:"column:content"`
	CreateDate string `json:"create_date,omitempty" gorm:"column:createdate"`
}

type User struct {
	Id            int64  `json:"id,omitempty" gorm:"primarykey"`
	Name          string `json:"name,omitempty" gorm:"column:name"`
	Password      string `gorm:"column:password"`
	FollowCount   int64  `json:"follow_count,omitempty" gorm:"column:followcount"`
	FollowerCount int64  `json:"follower_count,omitempty" gorm:"column:followercount"`
	IsFollow      bool   `json:"is_follow,omitempty" gorm:"-"`
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

func (User) TableName() string {
	return "user"
}

func (Video) TableName() string {
	return "video"
}

func (Comment) TableName() string {
	return "comment"
}

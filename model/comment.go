package model

import "time"

type Comment struct {
	Cid          int       `json:"cid" db:"cid"`
	AuthorId     int       `json:"author_id" db:"author_id"`
	PostId       int64     `json:"post_id" db:"post_id" binding:"required"`
	ParentId     int64     `json:"parent_id" db:"parent_id"`
	RootId       int64     `json:"root_id" db:"root_id"`
	CommentedUid int64     `json:"commented_uid" db:"commented_uid"`
	Content      string    `json:"content" db:"content" binding:"required"`
	CreateTime   time.Time `json:"create_time" db:"create_time"`
}

type ApiComment struct {
	Author   string `json:"author"`
	Stars    int64  `json:"stars"`
	ReplyNum int64  `json:"reply_num"`
	Started  bool   `json:"started"`
	*Comment
}

type ApiReply struct {
	Author          string `json:"author"`
	CommentedPeople string `json:"commented_people"`
	Stars           int64  `json:"stars"`
	Started         bool   `json:"started"`
	*Comment
}

type ApiCommentList struct {
	TotalNum int64         `json:"total_num"`
	Comments []*ApiComment `json:"comments"`
}

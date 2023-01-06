package model

import "time"

type Post struct {
	Type       int8      `json:"type" db:"type" binding:"required"`
	Pid        int       `json:"qid" db:"qid"`
	AuthorID   int       `json:"author_id" db:"author_id"`
	TopicID    int       `json:"topic_id" db:"topic_id" binding:"required"`
	Title      string    `json:"title" db:"title" binding:"required"`
	Content    string    `json:"content" db:"content" binding:"required"`
	CreateTime time.Time `json:"create_time" db:"create_time"`
	UpdateTime time.Time `json:"update_time" db:"update_time"`
}

type PostDetail struct {
	AuthorName string `json:"author_name"`
	LikeCount  int    `json:"like_count"`
	*Post
	*TopicDetail `json:"topic"`
}

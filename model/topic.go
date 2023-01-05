package model

import "time"

type Topic struct {
	Tid  int    `json:"topic-id" db:"tid"`
	Name string `json:"topic-name" db:"topic_name"`
}

type TopicDetail struct {
	Tid          int       `json:"topic_id" db:"tid"`
	Name         string    `json:"topic_name" db:"topic_name"`
	Introduction string    `json:"introduction" db:"introduction"`
	CreateTime   time.Time `json:"create-time" db:"create_time"`
}

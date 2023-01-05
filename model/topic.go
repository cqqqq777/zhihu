package model

type Topic struct {
	Tid  int    `json:"topic-id" db:"tid"`
	Name string `json:"topic-name" db:"topic_name"`
}

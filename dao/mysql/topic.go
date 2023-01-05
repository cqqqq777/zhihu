package mysql

import (
	"database/sql"
	g "zhihu/global"
	"zhihu/model"
)

const (
	GetAllTopicStr = "select tid , topic_name from topics "
)

func GetAllTopic() (data []*model.Topic, err error) {
	err = g.Mdb.Select(&data, GetAllTopicStr)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

package mysql

import (
	"database/sql"
	g "zhihu/global"
	"zhihu/model"
)

const (
	GetAllTopicStr = "select tid , topic_name from topics "
	TopicDetailStr = "select tid,topic_name,introduction,create_time from topics where tid = ?"
)

func GetAllTopic() (data []*model.Topic, err error) {
	err = g.Mdb.Select(&data, GetAllTopicStr)
	if err == sql.ErrNoRows {
		err = nil
	}
	return
}

func TopicDetails(tid int64) (data *model.TopicDetail, err error) {
	data = new(model.TopicDetail)
	err = g.Mdb.Get(data, TopicDetailStr, tid)
	if err == sql.ErrNoRows {
		err = ErrorInvalidId
		return nil, err
	}
	return
}

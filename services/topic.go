package services

import (
	"zhihu/dao/mysql"
	"zhihu/model"
)

func GetAllTopic() ([]*model.Topic, error) {
	return mysql.GetAllTopic()
}

func TopicDetail(tid int64) (*model.TopicDetail, error) {
	return mysql.TopicDetails(tid)
}

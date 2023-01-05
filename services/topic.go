package services

import (
	"zhihu/dao/mysql"
	"zhihu/model"
)

func GetAllTopic() ([]*model.Topic, error) {
	return mysql.GetAllTopic()
}

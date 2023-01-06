package services

import (
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
	"zhihu/model"
)

func GetAllTopic() (data []*model.Topic, err error) {
	//先从redis中查询，如果没有就在mysql中查询，查询后将数据保存到redis中
	val, err := redisdao.GetAllTopic()
	if err != nil {
		if err == redis.Nil {
			data, err = mysql.GetAllTopic()
			value, _ := json.Marshal(data)
			redisdao.SetAllTopic(value)
			return
		}
		return nil, err
	}
	if err = json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}
	return
}

func TopicDetail(tid int64) (data *model.TopicDetail, err error) {
	val, err := redisdao.GetTopicDetail(tid)
	if err != nil {
		if err == redis.Nil {
			data, err = mysql.TopicDetails(tid)
			if err == mysql.ErrorInvalidId {
				return
			}
			value, _ := json.Marshal(data)
			redisdao.SetTopicDetail(tid, value)
			return
		}
		return nil, err
	}
	if err = json.Unmarshal([]byte(val), &data); err != nil {
		return nil, err
	}
	return
}

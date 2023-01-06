package services

import (
	"encoding/json"
	"github.com/go-redis/redis/v9"
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
	"zhihu/model"
	"zhihu/utils"
)

func CreatePost(post *model.Post) error {
	pid, err := utils.GetID()
	if err != nil {
		return err
	}
	post.Pid = pid
	if err = mysql.CheckQuestion(post.Title); err != nil {
		return err
	}
	if err = mysql.CreatePost(post); err != nil {
		return err
	}
	return nil
}

func PostDetail(pid int64) (data *model.PostDetail, err error) {
	dataStr, err1 := redisdao.PostDetail(pid)
	if err1 == redis.Nil {
		data, err = mysql.PostDetail(pid)
		if err != nil {
			return nil, err
		}
		username, err2 := mysql.FindUsernameByUid(data.AuthorID)
		if err2 != nil {
			return nil, err2
		}
		data.AuthorName = username
		topic, err3 := mysql.TopicDetails(int64(data.TopicID))
		if err3 != nil {
			return nil, err3
		}
		data.TopicDetail = topic
		value, _ := json.Marshal(data)
		redisdao.SetTopicDetail(int64(data.Pid), value)
		return data, nil
	}
	if err = json.Unmarshal([]byte(dataStr), &data); err != nil {
		return nil, err
	}
	return
}

package services

import (
	"zhihu/dao/mysql"
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

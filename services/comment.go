package services

import (
	"zhihu/dao/mysql"
	"zhihu/model"
)

func CommentPost(comment *model.Comment) error {
	return mysql.CommentPost(comment)
}

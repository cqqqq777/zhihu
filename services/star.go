package services

import (
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
)

func StarPost(pid, uid int64) error {
	id, err := mysql.GetIdByUid(uid)
	if err != nil {
		return err
	}
	status, err := redisdao.GetUserStarStatus(pid, id)
	if err != nil {
		return err
	}
	switch status {
	case 1:
		err = redisdao.CancelStarPost(pid, id)
	case 0:
		err = redisdao.StarPost(pid, id)
	}
	return err
}

func StarComment(cid, uid int64) error {
	id, err := mysql.GetIdByUid(uid)
	if err != nil {
		return err
	}
	status, err := redisdao.GetUserStarCommentStatus(cid, id)
	if err != nil {
		return err
	}
	switch status {
	case 0:
		err = redisdao.StarComment(cid, id)
	case 1:
		err = redisdao.CancelStarComment(cid, id)
	}
	return err
}

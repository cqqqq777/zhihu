package services

import (
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
	"zhihu/model"
)

func CommentPost(comment *model.Comment) error {
	return mysql.CommentPost(comment)
}

func PostCommentList(pid, uid, page, size int64) (data *model.ApiCommentList, err error) {
	comments, err := mysql.PostCommentList(pid, page, size)
	if err != nil {
		return nil, err
	}
	data = new(model.ApiCommentList)
	data.Comments = make([]*model.ApiComment, 0)
	for _, comment := range comments {
		comment.Author, err = mysql.FindUsernameByUid(comment.AuthorId)
		if err != nil {
			continue
		}
		comment.ReplyNum, err = mysql.GetReplyNum(comment.Cid, pid)
		if err != nil {
			continue
		}
		comment.Stars, err = redisdao.GetCommentsStars(int64(comment.Cid))
		if err != nil {
			comment.Stars = 0
		}
		id, err := mysql.GetIdByUid(uid)
		if err != nil {
			comment.Started = false
		}
		status, err := redisdao.GetUserStarCommentStatus(int64(comment.Cid), id)
		if err != nil {
			comment.Started = false
		}
		switch status {
		case 0:
			comment.Started = false
		case 1:
			comment.Started = true
		}
		data.Comments = append(data.Comments, comment)
	}
	data.TotalNum, err = mysql.GetPostCommentsTotalNum(pid)
	if err != nil {
		return nil, err
	}
	return
}

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

func ReplyComment(comment *model.Comment) error {
	return mysql.ReplyComment(comment)
}

func ReplyList(cid, uid, page, size int64) (data []*model.ApiReply, err error) {
	replies, err := mysql.ReplyList(cid, page, size)
	if err != nil {
		return nil, err
	}
	data = make([]*model.ApiReply, 0)
	for _, reply := range replies {
		reply.Author, err = mysql.FindUsernameByUid(reply.AuthorId)
		if err != nil {
			continue
		}
		reply.CommentedPeople, err = mysql.FindUsernameByUid(int(reply.CommentedUid))
		if err != nil {
			continue
		}
		reply.Stars, _ = redisdao.GetCommentsStars(int64(reply.Cid))
		id, err := mysql.GetIdByUid(uid)
		if err != nil {
			reply.Started = false
		}
		status, err := redisdao.GetUserStarCommentStatus(int64(reply.Cid), id)
		if err != nil {
			reply.Started = false
		}
		switch status {
		case 0:
			reply.Started = false
		case 1:
			reply.Started = true
		}
		data = append(data, reply)
	}
	return data, nil
}

func UserReplies(uid int, page, size int64) (data *model.ApiUserReplies, err error) {
	replies, err := mysql.GetUserReplies(uid, page, size)
	if err != nil {
		return nil, err
	}
	data = new(model.ApiUserReplies)
	data.Replies = make([]*model.UserReplies, 0)
	for _, reply := range replies {
		reply.Author, err = mysql.FindUsernameByUid(reply.AuthorId)
		if err != nil {
			continue
		}
		data.Replies = append(data.Replies, reply)
	}
	data.TotalNum, err = mysql.GetUserReplyTotalNum(uid)
	if err != nil {
		return nil, err
	}
	return
}

func DeleteComment(cid, uid int64) error {
	authorId, err := mysql.GetAuthorIdByCid(cid)
	if err != nil {
		return err
	}
	if authorId != uid {
		return mysql.ErrorNoPermission
	}
	return mysql.DeleteComment(cid)
}

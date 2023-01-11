package mysql

import (
	g "zhihu/global"
	"zhihu/model"
)

const (
	CommentPostStr             = "insert into comments(cid,author_id,post_id,content,commented_uid) values(?,?,?,?,?)"
	PostCommentListStr         = "select * from comments where post_id = ? and root_id = 0 ORDER BY stars desc limit ?,?"
	GetPostCommentsTotalNumStr = "select count(cid) from comments where post_id =? and root_id = 0"
	GetReplyNumStr             = "select count(cid) from comments where post_id =? and root_id = ?"
	ReplyCommentStr            = "insert into comments(cid,author_id,post_id,parent_id,root_id,commented_uid,content) values(?,?,?,?,?,?,?) "
	ReplyListStr               = "select * from comments where root_id = ? ORDER BY stars desc limit ?,?"
	GetCidListStr              = "select cid from comments"
	SyncCommentStarsStr        = "update comments set stars = ? where cid = ?"
	GetUserRepliesStr          = "select cid,author_id,post_id,parent_id,root_id,commented_uid,content,create_time from comments where commented_uid = ? limit ?,?"
	GetUserReplyTotalNumStr    = "select count(cid) from comments where commented_uid = ?"
	GetAuthorIdByCidStr        = "select author_id from comments where cid = ?"
	DeleteCommentStr           = "delete from comments where cid = ?"
)

func CommentPost(comment *model.Comment) error {
	_, err := g.Mdb.Exec(CommentPostStr, comment.Cid, comment.AuthorId, comment.PostId, comment.Content, comment.CommentedUid)
	return err
}

func PostCommentList(pid, page, size int64) (comments []*model.ApiComment, err error) {
	comments = make([]*model.ApiComment, 0, 2)
	err = g.Mdb.Select(&comments, PostCommentListStr, pid, (page-1)*size, size)
	return
}

func GetReplyNum(cid int, pid int64) (replyNum int64, err error) {
	err = g.Mdb.Get(&replyNum, GetReplyNumStr, pid, cid)
	return
}

func GetPostCommentsTotalNum(pid int64) (totalNum int64, err error) {
	err = g.Mdb.Get(&totalNum, GetPostCommentsTotalNumStr, pid)
	return
}

func ReplyComment(comment *model.Comment) error {
	_, err := g.Mdb.Exec(ReplyCommentStr, comment.Cid, comment.AuthorId, comment.PostId, comment.ParentId, comment.RootId, comment.CommentedUid, comment.Content)
	return err
}

func ReplyList(cid, page, size int64) (replies []*model.ApiReply, err error) {
	replies = make([]*model.ApiReply, 0)
	err = g.Mdb.Select(&replies, ReplyListStr, cid, (page-1)*size, size)
	return
}

func GetCidList() (list []int64, err error) {
	list = make([]int64, 0)
	err = g.Mdb.Select(&list, GetCidListStr)
	return
}

func SyncCommentStars(cid, stars int64) error {
	_, err := g.Mdb.Exec(SyncCommentStarsStr, stars, cid)
	return err
}

func GetUserReplies(uid int, page, size int64) (replies []*model.UserReplies, err error) {
	replies = make([]*model.UserReplies, 0)
	err = g.Mdb.Select(&replies, GetUserRepliesStr, uid, (page-1)*size, size)
	return
}

func GetUserReplyTotalNum(uid int) (num int, err error) {
	err = g.Mdb.Get(&num, GetUserReplyTotalNumStr, uid)
	return
}

func GetAuthorIdByCid(cid int64) (authorId int64, err error) {
	err = g.Mdb.Get(&authorId, GetAuthorIdByCidStr, cid)
	return
}

func DeleteComment(cid int64) error {
	_, err := g.Mdb.Exec(DeleteCommentStr, cid)
	return err
}

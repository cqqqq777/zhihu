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

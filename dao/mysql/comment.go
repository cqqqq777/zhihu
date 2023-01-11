package mysql

import (
	g "zhihu/global"
	"zhihu/model"
)

const (
	CommentPostStr = "insert into comments(cid,author_id,post_id,content,commented_uid) values(?,?,?,?,?)"
)

func CommentPost(comment *model.Comment) error {
	_, err := g.Mdb.Exec(CommentPostStr, comment.Cid, comment.AuthorId, comment.PostId, comment.Content, comment.CommentedUid)
	return err
}

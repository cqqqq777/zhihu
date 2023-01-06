package mysql

import (
	"database/sql"
	g "zhihu/global"
	"zhihu/model"
)

const (
	CreatePostStr    = "insert into posts(pid,type,title,content,author_id,topic_id) values(?,?,?,?,?,?)"
	CheckQuestionStr = "select count(pid) from posts where title = ? and type = 1"
	PostDetailStr    = "select * from posts where pid = ?"
)

func CheckQuestion(title string) error {
	var count int8
	if err := g.Mdb.Get(&count, CheckQuestionStr, title); err != nil {
		return err
	}
	if count > 0 {
		return ErrorQuestionExist
	}
	return nil
}

func CreatePost(post *model.Post) (err error) {
	_, err = g.Mdb.Exec(CreatePostStr, post.Pid, post.Type, post.Title, post.Content, post.AuthorID, post.TopicID)
	return
}

func PostDetail(pid int64) (p *model.PostDetail, err error) {
	p = new(model.PostDetail)
	err = g.Mdb.Get(p, PostDetailStr, pid)
	if err == sql.ErrNoRows {
		return nil, ErrorInvalidId
	}
	return
}

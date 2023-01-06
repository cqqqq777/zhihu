package mysql

import (
	g "zhihu/global"
	"zhihu/model"
)

const (
	CreatePostStr    = "insert into posts(pid,type,title,content,author_id,topic_id) values(?,?,?,?,?,?)"
	CheckQuestionStr = "select count(pid) from posts where title = ? and type = 1"
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

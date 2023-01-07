package mysql

import (
	"database/sql"
	g "zhihu/global"
	"zhihu/model"
)

const (
	CreatePostStr              = "insert into posts(pid,type,title,content,author_id,topic_id) values(?,?,?,?,?,?)"
	CheckQuestionStr           = "select count(pid) from posts where title = ? and type = 1"
	PostDetailStr              = "select * from posts where pid = ?"
	QuestionListStr            = "select * from posts where  type = 1 ORDER BY create_time DESC limit ?,?"
	EssayListStr               = "select * from posts where  type = 2 ORDER BY create_time DESC limit ?,?"
	UserQuestionListStr        = "select * from posts where type = 1 and author_id = ? ORDER BY create_time DESC limit ?,?"
	UserEssayListStr           = "select * from posts where type = 2 and author_id = ? ORDER BY create_time DESC limit ?,?"
	GetQuestionTotalNumStr     = "select count(pid) from posts where type = 1 "
	GetEssayTotalNumStr        = "select count(pid) from posts where type = 2"
	GetUserQuestionTotalNumStr = "select count(pid) from posts where type = 1 and author_id = ?"
	GetUserEssayTotalNumStr    = "select count(pid) from posts where type = 2 and author_id = ?"
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

func QuestionList(page, size int64) (posts []*model.PostDetail, err error) {
	posts = make([]*model.PostDetail, 0, 2)
	err = g.Mdb.Select(&posts, QuestionListStr, (page-1)*size, size)
	return
}

func GetQuestionTotalNum() (num int, err error) {
	err = g.Mdb.Get(&num, GetQuestionTotalNumStr)
	return
}

func EssayList(page, size int64) (posts []*model.PostDetail, err error) {
	posts = make([]*model.PostDetail, 0, 2)
	err = g.Mdb.Select(&posts, EssayListStr, (page-1)*size, size)
	return
}

func GetEssayTotalNum() (num int, err error) {
	err = g.Mdb.Get(&num, GetEssayTotalNumStr)
	return
}
func UserQuestionList(page, size, uid int64) (posts []*model.PostDetail, err error) {
	posts = make([]*model.PostDetail, 0, 2)
	err = g.Mdb.Select(&posts, UserQuestionListStr, uid, (page-1)*size, size)
	return
}

func GetUserQuestionTotalNum(uid int64) (num int, err error) {
	err = g.Mdb.Get(&num, GetUserQuestionTotalNumStr, uid)
	return
}

func UserEssayList(page, size, uid int64) (posts []*model.PostDetail, err error) {
	posts = make([]*model.PostDetail, 0, 2)
	err = g.Mdb.Select(&posts, UserEssayListStr, uid, (page-1)*size, size)
	return
}

func GetUserEssayTotalNum(uid int64) (num int, err error) {
	err = g.Mdb.Get(&num, GetUserEssayTotalNumStr, uid)
	return
}

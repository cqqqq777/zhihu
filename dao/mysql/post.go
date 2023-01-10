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
	GetAuthorIdByPidStr        = "select author_id from posts where pid = ?"
	UpdatePostStr              = "update posts set type = ?,topic_id = ?,title = ?,content =? where pid =? "
	DeletePostStr              = "delete from posts where pid = ?"
	SearchPostsStr             = "select * from posts where title like ? ORDER BY create_time DESC limit ?,?"
	SearchTotalNumStr          = "select count(pid) from posts where  title like ? "
	GetPidListStr              = "select pid from posts"
	SyncPostStarsStr           = "update posts set stars = ? where pid =?"
	GetHotPostListStr          = "select * from posts ORDER BY stars desc limit ?,?"
	GetPostTotalNumStr         = "select count(pid) from posts "
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

func GetAuthorIdByPid(pid int) (uid int, err error) {
	err = g.Mdb.Get(&uid, GetAuthorIdByPidStr, pid)
	return
}

func UpdatePost(p *model.Post) error {
	_, err := g.Mdb.Exec(UpdatePostStr, p.Type, p.TopicID, p.Title, p.Content, p.Pid)
	return err
}

func DeletePost(pid int) error {
	_, err := g.Mdb.Exec(DeletePostStr, pid)
	return err
}

func SearchPosts(page, size int64, key string) (posts []*model.PostDetail, err error) {
	posts = make([]*model.PostDetail, 0, 2)
	err = g.Mdb.Select(&posts, SearchPostsStr, "%"+key+"%", (page-1)*size, size)
	return
}

func GetSearchPostTotalNum(key string) (num int, err error) {
	err = g.Mdb.Get(&num, SearchTotalNumStr, "%"+key+"%")
	return
}

func GetPidList() (pidList []int64, err error) {
	pidList = make([]int64, 0)
	err = g.Mdb.Select(&pidList, GetPidListStr)
	return
}

func SyncPostStars(pid, stars int64) error {
	_, err := g.Mdb.Exec(SyncPostStarsStr, stars, pid)
	return err
}

func GetHotPostList(page, size int64) (posts []*model.PostDetail, err error) {
	posts = make([]*model.PostDetail, 0)
	err = g.Mdb.Select(&posts, GetHotPostListStr, (page-1)*size, size)
	return
}

func GetPostTotalNum() (num int, err error) {
	err = g.Mdb.Get(&num, GetPostTotalNumStr)
	return
}

package services

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis/v9"
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
	g "zhihu/global"
	"zhihu/model"
	"zhihu/utils"
)

func CreatePost(post *model.Post) error {
	pid, err := utils.GetID()
	if err != nil {
		return err
	}
	post.Pid = pid
	if err = mysql.CheckQuestion(post.Title); err != nil {
		return err
	}
	if err = mysql.CreatePost(post); err != nil {
		return err
	}
	return nil
}

func PostDetail(pid int64) (data *model.PostDetail, err error) {
	//先从redis中获取数据，如果获取不到再到MySQL中获取数据，并在redis中设置缓存
	dataStr, err1 := redisdao.PostDetail(pid)
	if err1 == redis.Nil {
		data, err = mysql.PostDetail(pid)
		if err != nil {
			return nil, err
		}
		username, err2 := mysql.FindUsernameByUid(data.AuthorID)
		if err2 != nil {
			return nil, err2
		}
		data.AuthorName = username
		topic, err3 := mysql.TopicDetails(int64(data.TopicID))
		if err3 != nil {
			return nil, err3
		}
		data.TopicDetail = topic
		value, _ := json.Marshal(data)
		redisdao.SetPostDetail(int64(data.Pid), value)
		return data, nil
	}
	if err = json.Unmarshal([]byte(dataStr), &data); err != nil {
		return nil, err
	}
	return
}

func QuestionList(page, size int64) (data *model.ApiPostList, err error) {
	posts, err := mysql.QuestionList(page, size)
	if err != nil {
		return nil, err
	}
	data = new(model.ApiPostList)
	data.Posts = make([]*model.PostDetail, 0, len(posts))
	for _, post := range posts {
		username, err := mysql.FindUsernameByUid(post.AuthorID)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		topic, err := mysql.TopicDetails(int64(post.TopicID))
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		post.AuthorName = username
		post.TopicDetail = topic
		data.Posts = append(data.Posts, post)
	}
	var num int
	num, err = mysql.GetQuestionTotalNum()
	if err != nil {
		return nil, err
	}
	data.TotalNum = num
	return
}

func EssayList(page, size int64) (data *model.ApiPostList, err error) {
	posts, err := mysql.EssayList(page, size)
	if err != nil {
		g.Logger.Warn(fmt.Sprintf("%v", err))
		return nil, err
	}
	data = new(model.ApiPostList)
	data.Posts = make([]*model.PostDetail, 0, len(posts))
	for _, post := range posts {
		username, err := mysql.FindUsernameByUid(post.AuthorID)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		topic, err := mysql.TopicDetails(int64(post.TopicID))
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		post.AuthorName = username
		post.TopicDetail = topic
		data.Posts = append(data.Posts, post)
	}
	var num int
	num, err = mysql.GetEssayTotalNum()
	if err != nil {
		return nil, err
	}
	data.TotalNum = num
	return
}

func UserQuestionList(page, size, uid int64) (data *model.ApiPostList, err error) {
	posts, err := mysql.UserQuestionList(page, size, uid)
	if err != nil {
		return nil, err
	}
	data = new(model.ApiPostList)
	data.Posts = make([]*model.PostDetail, 0, len(posts))
	for _, post := range posts {
		username, err := mysql.FindUsernameByUid(post.AuthorID)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		topic, err := mysql.TopicDetails(int64(post.TopicID))
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		post.AuthorName = username
		post.TopicDetail = topic
		data.Posts = append(data.Posts, post)
	}
	var num int
	num, err = mysql.GetUserQuestionTotalNum(uid)
	if err != nil {
		return nil, err
	}
	data.TotalNum = num
	return
}

func UserEssayList(page, size, uid int64) (data *model.ApiPostList, err error) {
	posts, err := mysql.UserEssayList(page, size, uid)
	if err != nil {
		return nil, err
	}
	data = new(model.ApiPostList)
	data.Posts = make([]*model.PostDetail, 0, len(posts))
	for _, post := range posts {
		username, err := mysql.FindUsernameByUid(post.AuthorID)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		topic, err := mysql.TopicDetails(int64(post.TopicID))
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		post.AuthorName = username
		post.TopicDetail = topic
		data.Posts = append(data.Posts, post)
	}
	var num int
	num, err = mysql.GetUserEssayTotalNum(uid)
	if err != nil {
		return nil, err
	}
	data.TotalNum = num
	return
}

func UpdatePost(post *model.Post) error {
	authorId, err := mysql.GetAuthorIdByPid(post.Pid)
	if err != nil {
		return err
	}
	if authorId != post.AuthorID {
		return mysql.ErrorNoPermission
	}
	err = mysql.UpdatePost(post)
	redisdao.ClearPostCache(int64(post.Pid))
	return err
}

func DeletePost(uid, pid int) error {
	authorId, err := mysql.GetAuthorIdByPid(pid)
	if err != nil {
		return err
	}
	if authorId != uid {
		return mysql.ErrorNoPermission
	}
	err = mysql.DeletePost(pid)
	redisdao.ClearPostCache(int64(pid))
	return err
}

func SearchPosts(page, size int64, key string) (data *model.ApiPostList, err error) {
	posts, err := mysql.SearchPosts(page, size, key)
	if err != nil {
		return nil, err
	}
	data = new(model.ApiPostList)
	data.Posts = make([]*model.PostDetail, 0, len(posts))
	for _, post := range posts {
		username, err := mysql.FindUsernameByUid(post.AuthorID)
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		topic, err := mysql.TopicDetails(int64(post.TopicID))
		if err != nil {
			g.Logger.Warn(fmt.Sprintf("%v", err))
			continue
		}
		post.AuthorName = username
		post.TopicDetail = topic
		data.Posts = append(data.Posts, post)
	}
	var num int
	num, err = mysql.GetPostTotalNum(key)
	if err != nil {
		return nil, err
	}
	data.TotalNum = num
	return
}

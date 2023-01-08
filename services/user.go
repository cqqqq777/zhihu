package services

import (
	"crypto/tls"
	"errors"
	"fmt"
	"gopkg.in/gomail.v2"
	"math/rand"
	"time"
	"zhihu/dao/mysql"
	"zhihu/dao/redisdao"
	"zhihu/model"
	"zhihu/utils"
)

// PostVerification 发送验证码
func PostVerification(email string) error {
	vCode := rand.New(rand.NewSource(time.Now().UnixNano())).Int31n(1000000)
	message := fmt.Sprintf(`<div>
        <div>
            尊敬的用户，您好！
        </div>
        <div style="padding: 8px 40px 8px 50px;">
            <p>你本次的验证码为%06d,为了保证账号安全，验证码有效期为5分钟。请确认为本人操作，切勿向他人泄露，感谢您的理解与使用。</p>
        </div>
        <div>
            <p>此邮箱为系统邮箱，请勿回复。</p>
        </div>
    </div>`, vCode)

	// QQ 邮箱：
	// SMTP 服务器地址：smtp.qq.com（SSL协议端口：465/994 | 非SSL协议端口：25）
	// 163 邮箱：
	// SMTP 服务器地址：smtp.163.com（端口：25）
	host := "smtp.qq.com"
	port := 25
	userName := "498574842@qq.com"
	password := "aorwhfxnbzpobiai"

	m := gomail.NewMessage()
	m.SetHeader("From", userName) // 发件人
	// m.SetHeader("From", "alias"+"<"+userName+">") // 增加发件人别名

	m.SetHeader("To", email) // 收件人，可以多个收件人，但必须使用相同的 SMTP 连接
	//m.SetHeader("Cc", "******@qq.com")                  // 抄送，可以多个
	//m.SetHeader("Bcc", "******@qq.com")                 // 暗送，可以多个
	m.SetHeader("Subject", "Hello!") // 邮件主题

	// text/html 的意思是将文件的 content-type 设置为 text/html 的形式，浏览器在获取到这种文件时会自动调用html的解析器对文件进行相应的处理。
	// 可以通过 text/html 处理文本格式进行特殊处理，如换行、缩进、加粗等等
	m.SetBody("text/html", message)

	// text/plain的意思是将文件设置为纯文本的形式，浏览器在获取到这种文件时并不会对其进行处理
	// m.SetBody("text/plain", "纯文本")
	// m.Attach("test.sh")   // 附件文件，可以是文件，照片，视频等等
	// m.Attach("lolcatVideo.mp4") // 视频
	// m.Attach("lolcat.jpg") // 照片

	d := gomail.NewDialer(
		host,
		port,
		userName,
		password,
	)
	// 关闭SSL协议认证
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	if err != nil {
		return err
	}
	if err = redisdao.SetVerification(email, vCode); err != nil {
		return err
	}
	return nil
}

func Register(ParamUser *model.ParamRegisterUser) error {
	//检查用户名是否已存在
	if err := mysql.CheckUsername(ParamUser.Username); err != nil {
		return err
	}
	//检查邮箱是否已注册
	if err := mysql.CheckEmail(ParamUser.Email); err != nil {
		return err
	}
	//判断验证是否正确
	verification, err := redisdao.GetVerification(ParamUser.Email)
	if err != nil || verification != ParamUser.Verification {
		return redisdao.ErrorInvalidVerification
	}
	//获取uid
	uid, _ := utils.GetID()
	//添加新用户
	user := &model.User{
		Uid:      uid,
		Username: ParamUser.Username,
		//将密码加密存到数据库
		Password: utils.Md5(ParamUser.Password),
		Email:    ParamUser.Email,
	}
	return mysql.AddUser(user)
}

func LoginByUsername(user *model.ParamLoginUser) (int, string, error) {
	if err := mysql.CheckUsername(user.UsernameOrEmail); !errors.Is(err, mysql.ErrorUserExist) {
		if err == nil {
			return 0, "", mysql.ErrorUserNotExist
		}
		return 0, "", err
	}
	password, err := mysql.FindPasswordByUsername(user.UsernameOrEmail)
	if err != nil {
		return 0, "", err
	}
	if password != user.Password {
		return 0, "", mysql.ErrorWrongPassword
	}
	uid, err := mysql.FindUid(user.UsernameOrEmail)
	if err != nil {
		return 0, "", err
	}
	token, _ := utils.GenToken(uid)
	return uid, token, nil
}

func LoginByEmail(user *model.ParamLoginUser) (int, string, error) {
	if err := mysql.CheckEmail(user.UsernameOrEmail); !errors.Is(err, mysql.ErrorEmailExist) {
		if err == nil {
			return 0, "", mysql.ErrorEmailNotExist
		}
		return 0, "", err
	}
	password, err := mysql.FindPasswordByEmail(user.UsernameOrEmail)
	if err != nil {
		return 0, "", err
	}
	if password != user.Password {
		return 0, "", mysql.ErrorWrongPassword
	}
	uid, err := mysql.FindUid(user.UsernameOrEmail)
	if err != nil {
		return 0, "", err
	}
	token, _ := utils.GenToken(uid)
	return uid, token, nil
}

func RevisePassword(user *model.ParamReviseUser) error {
	password, err := mysql.FindPasswordByUid(user.Uid)
	if err != nil {
		return err
	}
	if user.OriPassword != password {
		return mysql.ErrorWrongPassword
	}
	if err = mysql.RevisePassword(user.NewPassword, user.Uid); err != nil {
		return err
	}
	return nil
}

func ReviseUsername(user *model.ParamReviseUser) error {
	if err := mysql.CheckUsername(user.NewUsername); err != nil {
		return err
	}
	if err := mysql.ReviseUsername(user.NewUsername, user.Uid); err != nil {
		return err
	}
	return nil
}

func ForgetPassword(user *model.ParamRegisterUser) error {
	if mysql.CheckEmail(user.Email) == nil {
		return mysql.ErrorEmailNotExist
	}
	verification, err := redisdao.GetVerification(user.Email)
	if err != nil || verification != user.Verification {
		return redisdao.ErrorInvalidVerification
	}
	uid, err := mysql.FindUid(user.Email)
	if err != nil {
		return err
	}
	if err = mysql.RevisePassword(user.Password, uid); err != nil {
		return err
	}
	return nil
}

func GetUserInfo(uid int64) (userInfo *model.User, err error) {
	userInfo, err = mysql.GetUserInfo(uid)
	return
}

func UpdateUserInfo(uid int, user *model.User) error {
	if uid != user.Uid {
		return mysql.ErrorNoPermission
	}
	err := mysql.UpdateUserInfo(user)
	return err
}

package mysql

import (
	g "zhihu/global"
	"zhihu/model"
)

const (
	FindUserByUsernameStr     = "select count(uid) from users where username = ?"
	FindUserByEmailStr        = "select count(uid) from users where email = ?"
	AddUserStr                = "insert into users(uid,username,password,email) values(?,?,?,?)"
	FindPasswordByUsernameStr = "select password from users where username = ?"
	FindPasswordByEmailStr    = "select password from users where email = ?"
	FindPasswordByUidStr      = "select password from users where uid = ?"
	FindUidStr                = "select uid from users where username = ? or email = ?"
	RevisePasswordStr         = "update users set password = ? where uid = ?"
	ReviseUsernameSte         = "update users set username = ? where uid =?"
)

func CheckUsername(username string) error {
	var count int64
	if err := g.Mdb.Get(&count, FindUserByUsernameStr, username); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

func CheckEmail(email string) error {
	var count int64
	if err := g.Mdb.Get(&count, FindUserByEmailStr, email); err != nil {
		return err
	}
	if count > 0 {
		return ErrorEmailExist
	}
	return nil
}

func AddUser(user *model.User) error {
	_, err := g.Mdb.Exec(AddUserStr, user.Uid, user.Username, user.Password, user.Email)
	return err
}

func FindPasswordByEmail(email string) (password string, err error) {
	if err := g.Mdb.Get(&password, FindPasswordByEmailStr, email); err != nil {
		return "", err
	}
	return
}

func FindPasswordByUsername(username string) (password string, err error) {
	if err := g.Mdb.Get(&password, FindPasswordByUsernameStr, username); err != nil {
		return "", err
	}
	return
}

func FindPasswordByUid(uid int) (password string, err error) {
	if err := g.Mdb.Get(&password, FindPasswordByUidStr, uid); err != nil {
		return "", err
	}
	return
}

func FindUid(UsernameOrEmail string) (uid int, err error) {
	err = g.Mdb.Get(&uid, FindUidStr, UsernameOrEmail, UsernameOrEmail)
	return
}

func RevisePassword(password string, uid int) error {
	_, err := g.Mdb.Exec(RevisePasswordStr, password, uid)
	return err
}

func ReviseUsername(username string, uid int) error {
	_, err := g.Mdb.Exec(ReviseUsernameSte, username, uid)
	return err
}

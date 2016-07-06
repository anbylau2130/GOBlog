package lib

import (
	"blog/models"
	"crypto/md5"
	"encoding/hex"
	"errors"
	"strconv"
)

//Strtomd5 生成 md5
func Strtomd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	rs := hex.EncodeToString(h.Sum(nil))
	return rs
}

//Pwdhash 密码加密
func Pwdhash(str string) string {
	return Strtomd5(str)
}

//StringsToJson 字符串转JSON
func StringsToJson(str string) string {
	rs := []rune(str)
	jsons := ""
	for _, r := range rs {
		rint := int(r)
		if rint < 128 {
			jsons += string(r)
		} else {
			jsons += "\\u" + strconv.FormatInt(int64(rint), 16) // json
		}
	}

	return jsons
}

// GetPassWord 计算用户密码
func GetPassWord(loginName string, password string) (newPassword string) {
	newPassword = loginName + "{{" + password + "}}"
	return newPassword
}

//CheckLogin 检查用户登录
func CheckLogin(username string, password string) (user models.SysOperator, err error) {

	user = models.GetOperatorByLoginName(username)
	if user.ID == 0 && user.LoginName != "admin" {
		return user, errors.New("用户不存在")
	}
	md5Pwd := Pwdhash(GetPassWord(username, password))
	if user.Password != md5Pwd {
		return user, errors.New("密码错误")
	}
	return user, nil
}

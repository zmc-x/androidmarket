package mall

import (
	"androidmarket/global"
	"androidmarket/model/mall"
	"androidmarket/model/mall/request"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
	"time"
)

type Malluser struct{}

// Userlogin 用户登录
func (m *Malluser) Userlogin(params request.Login) (bool, string) {
	temp := mall.User{}
	res := global.GlobalDB.Where("username = ? and password = ?", params.Username, params.Password).Find(&temp)
	if res.RowsAffected == 0 {
		// 登录失败
		return false, ""
	} else {
		// 登录成功 return the token
		s := setToken(temp)
		return true, s
	}
}

// Usersignup 注册
func (m *Malluser) Usersignup(params request.Login) (bool, error) {
	temp := mall.User{
		Uid:       strconv.Itoa(int(time.Now().Unix())),
		Username:  params.Username,
		Password:  params.Password,
		Privilege: 0,
	}
	res := global.GlobalDB.Create(&temp)
	if res.Error != nil {
		// 用户名已经重复了
		return false, res.Error
	} else {
		// 注册成功
		return true, nil
	}
}

// Updatepass 修改用户密码
func (m *Malluser) Updatepass(params request.Updatepass, uid string) (bool, error) {
	res := global.GlobalDB.Model(&mall.User{}).Where("uid = ?", uid).Update("password", params.Newpassword)
	// error不为nil
	if res.Error != nil {
		return false, res.Error
	} else {
		return true, nil
	}
}

// setToken 生成token
func setToken(user mall.User) string {
	mykey := []byte("hellothisishellozmc")
	type Myclaim struct {
		Uid string
		jwt.RegisteredClaims
	}

	m := Myclaim{
		Uid: user.Uid,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    "hellozmc",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &m)
	s, _ := token.SignedString(mykey)
	return s
}

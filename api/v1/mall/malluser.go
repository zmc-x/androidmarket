package mall

import (
	"androidmarket/middleware"
	mallresp "androidmarket/model/common/response"
	"androidmarket/model/mall/request"
	"github.com/gin-gonic/gin"
)

type Malluser struct{}

// Userlogin 用户登录
func (m Malluser) Userlogin(c *gin.Context) {
	// 获取用户的输入的信息
	logininfo := request.Login{}
	err := c.ShouldBindJSON(&logininfo)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if status, token := Userlogin(logininfo); status {
		mallresp.OkWithDetail(c, map[string]interface{}{"token": token}, "登录成功")
	} else {
		mallresp.FailWithDetail(c, nil, "登录失败，用户名or密码错误！")
	}
}

// Usersignup 用户注册
func (m Malluser) Usersignup(c *gin.Context) {
	// 获取用户的输入的信息
	signupinfo := request.Login{}
	err := c.ShouldBindJSON(&signupinfo)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if status, err := UserSignup(signupinfo); status {
		mallresp.OkWithDetail(c, nil, "恭喜你，注册成功！")
	} else {
		mallresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, "该用户名已经存在！")
	}
}

// Userupdatepass 修改密码
func (m Malluser) Userupdatepass(c *gin.Context) {
	// 获取新密码
	newpass := request.Updatepass{}
	err := c.ShouldBindJSON(&newpass)
	if err != nil {
		mallresp.FailParam(c)
		return
	}
	if status, err := UserUpdatePass(newpass, middleware.Uid); status {
		mallresp.OkWithDetail(c, nil, "用户密码修改成功！")
	} else {
		if err != nil {
			mallresp.FailWithDetail(c, map[string]interface{}{"error": err.Error()}, "用户密码修改失败！")
		} else {
			mallresp.FailWithDetail(c, nil, "用户密码修改失败！")
		}
	}
}

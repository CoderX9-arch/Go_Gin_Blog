package user_api

import (
	"blog/global"
	"blog/models"
	"blog/models/res"
	"blog/utils/jwts"
	"blog/utils/pwd"
	"github.com/gin-gonic/gin"
)

type EmailLoginRequest struct {
	UserName string `json:"username" binding:"required" msg:"请输入用户名"` //用户名
	Password string `json:"password" binding:"required" msg:"请输入用户密码"`
}

func (UserApi) EmailLoginView(c *gin.Context) {
	var cr EmailLoginRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithErr(err, &cr, c)
		return
	}
	var userModel models.UserModel
	err = global.DB.Take(&userModel, "user_name=?", cr.UserName).Error
	if err != nil {
		global.Log.Warn("用户不存在", c)
		res.FailWithMsg("用户密码错误", c)
		return
	}
	//校验密码
	//fmt.Println(userModel.Password)
	//fmt.Println(cr.Password)

	isCheck := pwd.CheckPwd(userModel.Password, cr.Password)
	if !isCheck {
		global.Log.Warn("用户密码错误")
		res.FailWithMsg("密码错误", c)
		return
	}
	//登录成功，生成token
	token, err := jwts.GenToken(jwts.JwtPayLoad{
		Nickname: userModel.NickName,
		Role:     int(userModel.Role),
		UserID:   string(userModel.ID),
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMsg("token生成失败", c)
		return
	}
	res.OkWithData(token, c)
}

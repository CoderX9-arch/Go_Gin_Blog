package user_api

import (
	"blog/models"
	"blog/models/ctype"
	"blog/models/res"
	"blog/service/common"
	"blog/utils/desens"
	"blog/utils/jwts"
	"github.com/gin-gonic/gin"
)

type UserResponse struct {
}

func (UserApi) UserListView(c *gin.Context) {
	//判断是否是管理员
	token := c.Request.Header.Get("token")
	if token == "" {
		res.FailWithMsg("未携带token", c)
		return
	}
	claims, err := jwts.ParseToken(token)
	if err != nil {
		res.FailWithMsg("token错误", c)
		return
	}

	var page models.PageInfo

	if err := c.ShouldBind(&page); err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, _ := common.ComList(models.UserModel{}, common.Option{
		PageInfo: page,
	})
	var users []models.UserModel

	for _, user := range list {

		if ctype.Role(claims.Role) != ctype.PermissionUser {
			user.UserName = ""
		}
		user.Tel = desens.DesensitizatiionTel(user.Tel)
		user.Email = desens.DesensitizatiionEmail(user.Email)
		users = append(users, user)

	}

	res.OkWithList(count, users, c)
}

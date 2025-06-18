package flag

import (
	"blog/global"
	"blog/models"
	"blog/models/ctype"
	"blog/utils/pwd"
	"fmt"
)

func CreateUser(permission string) {
	var (
		userName   string
		nickName   string
		passWord   string
		rePassWord string
		email      string
	)
	fmt.Println("请输入用户名")
	fmt.Scan(&userName)
	fmt.Println("请输入昵称")
	fmt.Scan(&nickName)
	fmt.Println("请输入邮箱")
	fmt.Scan(&email)
	fmt.Println("请输入密码")
	fmt.Scan(&passWord)
	fmt.Println("请再次输入密码")
	fmt.Scan(&rePassWord)

	var count int64
	var userModel models.UserModel
	global.DB.Take(&userModel, "user_name = ?", userName).Count(&count)
	//fmt.Println(err)
	if count > 0 {
		global.Log.Error("用户名已存在，请重新输入")
		return
	}
	if passWord != rePassWord {
		global.Log.Error("密码不一致，请重新输入")
		return
	}
	role := ctype.PermissionUser
	if permission == "admin" {
		role = ctype.Permissionadmin
	}
	avater := "/uploads/avatar/default.jpg"

	hashPwd := pwd.HsahPwd(passWord)
	err := global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashPwd,
		Email:      email,
		Role:       role,
		Avatar:     avater,
		IP:         "127.0.0.1",
		Addr:       "内网地址",
		SignStatus: ctype.SignEmail,
	}).Error
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功", userName)

}

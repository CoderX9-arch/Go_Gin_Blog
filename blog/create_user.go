package main

import (
	"blog/global"
	"blog/models"
	"blog/models/ctype"
	"blog/utils/pwd"
	"fmt"
)

func CreateUser(permissions string) {
	var (
		userName   string
		nickName   string
		password   string
		rePassword string
		email      string
	)

	fmt.Printf("创建用户，请输入用户名")
	fmt.Scan(&userName)
	fmt.Printf("创建用户，请输入昵称")
	fmt.Scan(&nickName)
	fmt.Printf("创建用户，请输入密码")
	fmt.Scan(&password)
	fmt.Printf("创建用户，请再次输入密码")
	fmt.Scan(&rePassword)
	fmt.Printf("创建用户，请输入邮箱")
	fmt.Scan(&email)

	var userModel models.UserModel
	err := global.DB.Take(&userModel, "user_name=?", userName)
	if err != nil {
		global.Log.Error("用户名已存在，请重新输入")
		return
	}
	if password != password {
		global.Log.Error("两次密码，请重新输入")
		return
	}
	//密码加密
	hashpwd := pwd.HsahPwd(password)

	role := ctype.PermissionUser
	err = global.DB.Create(&models.UserModel{
		NickName:   nickName,
		UserName:   userName,
		Password:   hashpwd,
		Email:      email,
		Role:       role,
		SignStatus: ctype.SignEmail,
	})
	fmt.Println(err)
	if err != nil {
		global.Log.Error(err)
		return
	}
	global.Log.Infof("用户%s创建成功", userName)

}

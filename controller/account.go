package controller

import (
	"github.com/kataras/iris/v12"
	"irisweb/provider"
	"irisweb/middware"
	"irisweb/model"
	"time"
	"fmt"
)

func Login(ctx iris.Context) {
	username := ctx.PostValue("username") 
	password := ctx.PostValue("password")
	fmt.Printf(username)
	var db = provider.GetDefaultDB()
	var users []model.User
	db.Find(&users, "user_name = ?", username)
	if len(users) == 0 {
		ctx.WriteString("用户不存在")
	} else {
		if users[0].UserName == username && users[0].Password == password {
			token := middware.GenerateToken(users[0].Id)
    		fmt.Println("生成JWT token:", token)
			ctx.JSON(iris.Map{
				"code": 200,
				"data": token,
				"msg":  "登录成功",
			})
		} else {
			ctx.WriteString("账户或密码错误")
		}
	}
	
	
}
func Register(ctx iris.Context) {
	username := ctx.PostValue("username") 
	password := ctx.PostValue("password")
	var db = provider.GetDefaultDB()
	var users []model.User
	var user model.User
	db.Find(&users, "user_name = ?", username)
	if len(users) == 0 {
		// 数据不存在
		user.UserName = username
		user.Password = password
		user.CreatedTime = time.Now().Unix()
		result := db.Create(&user)
		if result.Error != nil {
			ctx.WriteString("注册失败")
			return
		}
		ctx.WriteString("注册成功")
	} else {
		// 数据存在
		ctx.WriteString("用户已存在")
	}
	
	
}
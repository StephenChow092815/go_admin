package controller

import (
	"fmt"
	"irisweb/model"
	"irisweb/provider"

	"github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
)

func GetUsers(ctx iris.Context) {
	var db = provider.GetDefaultDB()
	var users []model.User
	result := db.Find(&users)
	if result.Error != nil {
		fmt.Printf("查询数据失败：%v", result.Error)
	}
	ctx.JSON(iris.Map{
		"code": 200,
		"data": users,
		"msg":  "Success",
	})
}
func GetUserInfo(ctx iris.Context) {
	fmt.Printf("jinru：%n", 1)
	jwtInfo := ctx.Values().Get("jwt").(*jwt.Token)
	id := jwtInfo.Claims.(jwt.MapClaims)["Uid"].(float64)
	var user model.User
	var db = provider.GetDefaultDB()
	result := db.Find(&user, "id = ?", id)
	if result.Error != nil {
		ctx.JSON(iris.Map{
			"code": 500,
			"data": nil,
			"msg":  "系统错误",
		})
	} else {
		ctx.JSON(iris.Map{
			"code": 200,
			"data": user,
			"msg":  "Success",
		})
	}
}

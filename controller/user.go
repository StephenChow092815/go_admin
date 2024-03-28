package controller

import (
	"github.com/kataras/iris/v12"
	"irisweb/provider"
	"irisweb/model"
	"fmt"
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
package controller


import (
	"github.com/kataras/iris/v12"
	"irisweb/provider"
	"irisweb/model"
	"time"
	"fmt"
	"strconv"
)

// 查询分类列表
func GetGoodsList(ctx iris.Context) {
	var db = provider.GetDefaultDB()
	var goodss []model.Goods
	
	result := db.Find(&goodss)
	if result.Error != nil {
        fmt.Printf("查询数据失败：%v", result.Error)
    }
	ctx.JSON(iris.Map{
		"code": 200,
		"data": goodss,
		"msg":  "Success",
	})
}
// 新增分类
func AddGoods(ctx iris.Context) {
	goodsName := ctx.PostValue("name") 
	var db = provider.GetDefaultDB()
	var goodss []model.Goods
	var goods model.Goods
	db.Find(&goodss, "name = ?", goodsName)
	if len(goodss) == 0 {
		// 数据不存在
		goods.Name = goodsName
		goods.Description = ctx.PostValue("description") 
		goods.Price = ctx.PostValue("price") 
		goods.Url = ctx.PostValue("url")
		id,_ := strconv.Atoi(ctx.PostValue("categoryId"))
		goods.CategoryId = id
		goods.CreatedTime = time.Now().Unix()
		result := db.Create(&goods)
		if result.Error != nil {
			ctx.JSON(iris.Map{
				"code": 500,
				"data": nil,
				"msg":  "创建商品失败",
			})
			return
		}
		ctx.JSON(iris.Map{
			"code": 200,
			"data": nil,
			"msg":  "创建成功",
		})
	} else {
		// 数据存在
		ctx.JSON(iris.Map{
			"code": 500,
			"data": nil,
			"msg":  "商品已存在",
		})
	}
}
// 编辑分类
func EditGoods(ctx iris.Context) {
	id := ctx.PostValue("id") 
	var db = provider.GetDefaultDB()
	var goodss []model.Goods
	db.Find(&goodss, "id = ?", id)
	if len(goodss) != 0 {
		var goods model.Goods

		db.Where("id = ?", id).Take(&goods)

		data := make(map[string]interface{})
		data["name"] = ctx.PostValue("name")  //零值字段
		data["description"] = ctx.PostValue("description")
		data["price"] = ctx.PostValue("price")
		data["url"] = ctx.PostValue("url")
		category_id,_ := strconv.Atoi(ctx.PostValue("categoryId"))
		data["category_id"] = category_id
		data["updated_time"] = time.Now().Unix()
		
		result := db.Model(&goods).Updates(data)
		if result.Error != nil {
			ctx.JSON(iris.Map{
				"code": 500,
				"data": nil,
				"msg":  "编辑商品失败",
			})
			return
		}
		ctx.JSON(iris.Map{
			"code": 200,
			"data": nil,
			"msg":  "编辑成功",
		})
	} else {
		// 数据存在
		ctx.JSON(iris.Map{
			"code": 500,
			"data": nil,
			"msg":  "商品不存在",
		})
	}
}
// 删除分类
func DeleteGoods(ctx iris.Context) {
	id := ctx.PostValue("id")
	var db = provider.GetDefaultDB()
	var goodss []model.Goods
	db.Find(&goodss, "id = ?", id)
	if len(goodss) != 0 {
		var goods model.Goods

		db.Where("id = ?", id).Take(&goods)

		
		
		result := db.Delete(&goods)
		if result.Error != nil {
			ctx.JSON(iris.Map{
				"code": 500,
				"data": nil,
				"msg":  "删除商品失败",
			})
			return
		}
		ctx.JSON(iris.Map{
			"code": 200,
			"data": nil,
			"msg":  "删除成功",
		})
	} else {
		// 数据存在
		ctx.JSON(iris.Map{
			"code": 500,
			"data": nil,
			"msg":  "商品不存在",
		})
	}
}
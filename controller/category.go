package controller

import (
	"fmt"
	"irisweb/model"
	"irisweb/provider"
	"time"

	"github.com/kataras/iris/v12"
)

// 查询分类列表
func GetCategoryList(ctx iris.Context) {
	var db = provider.GetDefaultDB()
	var categorys []model.Category

	result := db.Find(&categorys)
	if result.Error != nil {
		fmt.Printf("查询数据失败：%v", result.Error)
	}
	ctx.JSON(iris.Map{
		"code": 200,
		"data": categorys,
		"msg":  "Success",
	})
}

// 新增分类
func AddCategory(ctx iris.Context) {
	requestJSON := map[string]string{}
	if err := ctx.ReadJSON(&requestJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid JSON"})
	}
	categoryname := requestJSON["name"]
	tag := requestJSON["tag"]
	var db = provider.GetDefaultDB()
	var categorys []model.Category
	var category model.Category
	db.Find(&categorys, "name = ?", categoryname)
	if len(categorys) == 0 {
		// 数据不存在
		category.Name = categoryname
		category.Tag = tag
		category.CreatedTime = time.Now().Unix()
		result := db.Create(&category)
		if result.Error != nil {
			ctx.JSON(iris.Map{
				"code": 500,
				"data": nil,
				"msg":  "创建分类失败",
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
			"msg":  "分类已存在",
		})
	}
}

// 编辑分类
func EditCategory(ctx iris.Context) {
	requestJSON := make(map[string]interface{})
	if err := ctx.ReadJSON(&requestJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid JSON"})
	}
	categoryname := requestJSON["name"].(string)
	tag := requestJSON["tag"].(string)
	id := requestJSON["id"].(float64)
	var db = provider.GetDefaultDB()
	var categorys []model.Category

	db.Find(&categorys, "id = ?", id)
	if len(categorys) != 0 {
		var category model.Category

		db.Where("id = ?", id).Take(&category)

		data := make(map[string]interface{})
		data["name"] = categoryname //零值字段
		data["tag"] = tag
		data["updated_time"] = time.Now().Unix()

		result := db.Model(&category).Updates(data)
		if result.Error != nil {
			ctx.JSON(iris.Map{
				"code": 500,
				"data": nil,
				"msg":  "编辑分类失败",
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
			"msg":  "分类不存在",
		})
	}
}

// 删除分类
func DeleteCategory(ctx iris.Context) {
	requestJSON := map[string]float64{}
	if err := ctx.ReadJSON(&requestJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid JSON"})
	}
	id := requestJSON["id"]
	var db = provider.GetDefaultDB()
	var categorys []model.Category
	db.Find(&categorys, "id = ?", id)
	if len(categorys) != 0 {
		var category model.Category

		db.Where("id = ?", id).Take(&category)

		result := db.Delete(&category)
		if result.Error != nil {
			ctx.JSON(iris.Map{
				"code": 500,
				"data": nil,
				"msg":  "删除分类失败",
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
			"msg":  "分类不存在",
		})
	}
}

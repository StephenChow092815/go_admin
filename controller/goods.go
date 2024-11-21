package controller

import (
	"fmt"
	"irisweb/model"
	"irisweb/provider"
	"time"

	"github.com/kataras/iris/v12"
)

// 查询分类列表
func GetGoodsList(ctx iris.Context) {
	GetPagination(ctx, &[]model.Goods{})
}

func GetPagination(ctx iris.Context, modelInstance interface{}) {
	var db = provider.GetDefaultDB()
	requestJSON := make(map[string]interface{})
	if err := ctx.ReadJSON(&requestJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid JSON"})
	}

	var total int64

	current := int(requestJSON["current"].(float64))
	size := int(requestJSON["size"].(float64))

	db.Model(modelInstance).Count(&total)
	// 获取页数
	pages := int64(total / int64(size))
	if total%int64(size) != 0 {
		pages++
	}
	result := db.Limit(size).Offset((current - 1) * size).Find(modelInstance)
	if result.Error != nil {
		fmt.Println("查询数据失败：", result.Error)
	} else if result.RowsAffected >= 0 {
		fmt.Println("modelInstance：------>", modelInstance)
		ctx.JSON(iris.Map{
			"code": 200,
			"data": model.Pagination{
				Current: current,
				Size:    size,
				Total:   total,
				Pages:   pages,
				List:    modelInstance,
			},
			"msg": "Success",
		})
	}
}

// 新增分类
func AddGoods(ctx iris.Context) {
	requestJSON := make(map[string]interface{})
	if err := ctx.ReadJSON(&requestJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid JSON"})
	}
	goodsName := requestJSON["name"].(string)

	var db = provider.GetDefaultDB()
	var goodss []model.Goods
	var goods model.Goods
	db.Find(&goodss, "name = ?", goodsName)
	if len(goodss) == 0 {
		// 数据不存在
		goods.Name = goodsName
		goods.Description = requestJSON["description"].(string)
		goods.Price = requestJSON["price"].(string)
		goods.Url = requestJSON["url"].(string)
		categoryId := requestJSON["category_id"].(float64)
		// // id, _ := strconv.Atoi(categoryId)
		fmt.Printf("%s is an int: %d\n", "categoryId", categoryId)
		goods.CategoryId = categoryId
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
	requestJSON := make(map[string]interface{})
	if err := ctx.ReadJSON(&requestJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid JSON"})
	}
	id := requestJSON["id"].(float64)
	var db = provider.GetDefaultDB()
	var goodss []model.Goods
	db.Find(&goodss, "id = ?", id)
	if len(goodss) != 0 {
		var goods model.Goods

		db.Where("id = ?", id).Take(&goods)

		data := make(map[string]interface{})
		data["name"] = requestJSON["name"].(string) //零值字段
		data["description"] = requestJSON["description"].(string)
		data["price"] = requestJSON["price"].(string)
		data["url"] = requestJSON["url"].(string)
		category_id := requestJSON["category_id"].(float64)
		fmt.Printf("%s is an int: %d\n", "categoryId", category_id)
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
	requestJSON := map[string]float64{}
	if err := ctx.ReadJSON(&requestJSON); err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid JSON"})
	}
	id := requestJSON["id"]
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

package mockData

import (
	"fmt"
	"irisweb/model"
	"irisweb/provider"
	"strconv"
)

func InsertMockGoodsData() {
	fmt.Println("模拟开始！")
	var db = provider.GetDefaultDB()

	var goodss []model.Goods
	for i := 0; i < 100; i++ {
		goodss = append(goodss, model.Goods{
			Name:        fmt.Sprintf("商品%d", i),
			Description: fmt.Sprintf("商品描述%d", i),
			Price:       strconv.Itoa(100 + i),
			Url:         fmt.Sprintf("www.baidu%d.com", i),
			CategoryId:  0,
		})
	}

	// 开启事务
	tx := db.Begin()
	// 错误回滚
	defer tx.Rollback()

	result := tx.Create(goodss).Error
	if result != nil {
		fmt.Println("插入失败：", result)
	}

	// 提交事务
	tx.Commit()

	fmt.Println("数据在事务中批量插入成功！")
}

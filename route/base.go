package route

import (
	"github.com/kataras/iris/v12"
	"irisweb/controller"
	"irisweb/middware"
	"github.com/iris-contrib/middleware/jwt"
)



func Register(app *iris.Application) {
	j := jwt.New(jwt.Config{
		// Extractor属性可以选择从什么地方获取jwt进行验证，默认从http请求的header中的Authorization字段提取，也可指定为请求参数中的某个字段
	
		// 从请求参数token中提取
		// Extractor: jwt.FromParameter("token"),
	
		// 从请求头的Authorization字段中提取，这个是默认值
		Extractor: jwt.FromAuthHeader,
	
		// 设置一个函数返回秘钥，关键在于return []byte("这里设置秘钥")
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
		  return middware.SigKey, nil
		},
	
		// 设置一个加密方法
		SigningMethod: jwt.SigningMethodHS256,
	})
	app.HandleDir("/uploads", "./uploads")
	// 用户模块
	userhandler := app.Party("/user")
	//公共模块
	commonhandler := app.Party("/common")
	// 分类模块
	categoryhandler := app.Party("/category")
	// 商品模块
	goodshandler := app.Party("/goods")

	app.Post("/login", controller.Login)
	app.Post("/register", controller.Register)
	userhandler.Get("/",j.Serve, controller.GetUsers)
	commonhandler.Post("/upload",j.Serve, controller.UploadFile)

	// 分类模块
	categoryhandler.Get("/getCategoryList",j.Serve, controller.GetCategoryList)
	categoryhandler.Post("/",j.Serve, controller.AddCategory)
	categoryhandler.Put("/",j.Serve, controller.EditCategory)
	categoryhandler.Delete("/",j.Serve, controller.DeleteCategory)

	// 商品模块
	goodshandler.Get("/getGoodsList",j.Serve, controller.GetGoodsList)
	goodshandler.Post("/",j.Serve, controller.AddGoods)
	goodshandler.Put("/",j.Serve, controller.EditGoods)
	goodshandler.Delete("/",j.Serve, controller.DeleteGoods)
}
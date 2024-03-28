package controller

import (
	"github.com/kataras/iris/v12"
    "io"
    "os"
)
const maxSize = 8 * iris.MB
const loc = "http://localhost:4001"
func UploadFile(ctx iris.Context) {
	ctx.SetMaxRequestBodySize(maxSize)

        file, info, err := ctx.FormFile("file")
        if err != nil {
            iris.New().Logger().Info(err.Error())
            ctx.JSON(iris.Map{
                "code": 500,
                "data": nil,
                "msg":  "文件上传失败",
                
            })
            return
        }
        //最后要关闭文件
        defer file.Close()

        //获取文件名称`
        fname := info.Filename
        //把文件上传到哪里
        out, err := os.OpenFile("./uploads/"+fname, os.O_WRONLY|os.O_CREATE, 0666)
        if err != nil {
            iris.New().Logger().Info(err.Error())
            ctx.JSON(iris.Map{
                "code": 500,
                "data": nil,
                "msg":  "文件位置打开失败",
            })
            return
        }

        //我们打印一下文件路径
        iris.New().Logger().Info("文件路径："+out.Name())
        //最终也不要忘记关闭上传之后的文件流
        defer out.Close()

        //拷贝文件到指定位置
        _, err = io.Copy(out, file)
        if err != nil {
            ctx.JSON(iris.Map{
                "code": 500,
                "data": nil,
                "msg":  "文件上传彻底失败",
            })
            return
        }else {
            ctx.JSON(iris.Map{
                "code": 200,
                "data": loc + "/uploads/" + fname,
                "msg":  "文件上传成功",
            })
        }

}
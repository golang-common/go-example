// @Author: Perry
// @Date  : 2020/1/18
// @Desc  :
/*
首先swag init
然后运行main函数
反问http://127.0.0.1:8080/swagger/index
*/

package main

import (
	_ "dpy/exp/gin_exp/swagger_exp/docs"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

// @title 测试
// @version 0.0.1
// @description  测试
// @BasePath /api/v1/
func main() {
	r := gin.New()

	v1 := r.Group("/api/v1")
	v1.GET("/record/:userId", record)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	r.Run()
}

// @Summary 接口概要说明
// @Description 接口详细描述信息
// @Tags 用户信息
//swagger API分类标签, 同一个tag为一组
// @accept json
//浏览器可处理数据类型，浏览器默认发 Accept: */*
// @Produce  json
//设置返回数据的类型和编码
// @Param id path int true "ID"
//url参数：（name；参数类型[query(?id=),path(/123)]；数据类型；required；参数描述）
// @Param name query string false "name"
// @Success 200 {string} string "{"RetCode":0,"UserInfo":{},"Action":"GetAllUserResponse"}"
//成功返回的数据结构， 最后是示例
// @Failure 400 {string} string "{"RetCode":0,"UserInfo":{},"Action":"GetAllUserResponse"}"
// @Router /test/{id} [get]    //路由信息，一定要写上
func record(c *gin.Context) {
	c.String(http.StatusOK, "ok")
}

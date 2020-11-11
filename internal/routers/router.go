package routers

import (
	_ "github.com/catherine.li/go_blog/docs"
	"github.com/catherine.li/go_blog/global"
	"github.com/catherine.li/go_blog/internal/middleware"
	"github.com/catherine.li/go_blog/internal/routers/api"
	"github.com/catherine.li/go_blog/internal/routers/api/v1"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func NewRouter() *gin.Engine {

	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.Use(middleware.Translations())

	article := v1.NewArticle()
	tag := v1.NewTag()

	upload := api.NewUpload()
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload/file", upload.UploadFile)
	r.POST("/auth", api.GetAuth)

	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))
	apiv1 := r.Group("/api/v1")
	apiv1.Use() //middleware.JWT()
	{
		// 创建标签
		apiv1.POST("/tags", tag.Create)
		// 删除指定标签
		apiv1.DELETE("/tags/:id", tag.Delete)
		// 更新指定标签
		apiv1.PUT("/tags/:id", tag.Update)
		// 获取标签列表
		apiv1.GET("/tags", tag.List)

		// 创建文章
		apiv1.POST("/articles", article.Create)
		// 删除指定文章
		apiv1.DELETE("/articles/:id", article.Delete)
		// 更新指定文章
		apiv1.PUT("/articles/:id", article.Update)
		// 获取指定文章
		apiv1.GET("/articles/:id", article.Get)
		// 获取文章列表
		apiv1.GET("/articles", article.List)
	}

	return r
}

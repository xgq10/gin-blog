package routers

import (
	_ "gin-blog/docs"
	"gin-blog/middleware"
	"gin-blog/routers/api"
	"gin-blog/routers/api/v1"
	"gin-blog/utils/export"
	"gin-blog/utils/upload"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	gs "github.com/swaggo/gin-swagger"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.GET("/auth_service", v1.GetAuth)
	r.GET("/upload", api.UploadImage)
	r.GET("/swagger/*any", gs.WrapHandler(swaggerFiles.Handler))

	tag := r.Group("/api/v1")
	tag.Use(middleware.JWT())
	{
		//获取标签列表
		tag.GET("/tags", v1.GetTags)
		//新建标签
		tag.POST("/tags", v1.AddTag)
		//更新指定标签
		tag.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		tag.DELETE("/tags/:id", v1.DeleteTag)
		// 导出标签
		tag.POST("/tags/export", v1.ExportTag)
		// 导入标签
		tag.POST("/tags/import", v1.ImportTag)
	}

	article := r.Group("/api/v1")
	article.Use(middleware.JWT())
	{
		//获取文章列表
		article.GET("/articles", v1.GetArticles)
		//获取指定文章
		article.GET("/articles/:id", v1.GetArticle)
		//新建文章
		article.POST("/articles", v1.AddArticle)
		//更新指定文章
		article.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		article.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}

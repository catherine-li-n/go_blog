package v1

import (
	"github.com/catherine.li/go_blog/pkg/app"
	"github.com/catherine.li/go_blog/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "Hello world!",
	})
	app.NewResponse(c).ToErrorResponse(errcode.ServerError)
	return
}
func (a Article) List(c *gin.Context)   {}
func (a Article) Create(c *gin.Context) {}
func (a Article) Update(c *gin.Context) {}
func (a Article) Delete(c *gin.Context) {}

package router

import (
	"goblog/model"
	"goblog/model_view"
	"goblog/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListBlogs(c *gin.Context) {
	// offset and limit are in "offset=5&limit=15" form, so we should use Query instead of Param here
	offsetStr := c.Query("offset")
	limitStr := c.Query("limit")
	typeId := c.Query("type_id")
	tagId := c.Query("tag_id")
	pageOptions := model.ParsePageOptions(offsetStr, limitStr, 10, 0)
	code := http.StatusOK
	var blogs []model.Blog
	if typeId != "" {
		blogs = service.ListBlogsByTypeId(pageOptions, typeId)
	} else if tagId != "" {
		blogs = service.ListBlogsByTagId(pageOptions, tagId)
	} else {
		blogs = service.ListBlogs(pageOptions)
	}
	var blogsView []model_view.BlogView
	for _, blog := range blogs {
		blogType := service.GetTypeById(blog.TypeId)
		tags := service.ListTagsByBlogId(blog.BlogId)
		blogsView = append(blogsView, model_view.BlogView{
			Blog: blog,
			Type: blogType,
			Tags: tags,
		})
	}
	c.JSON(code, blogsView)
}

func GetBlog(c *gin.Context) {
	code := http.StatusOK
	blogId := c.Param("blog_id")

	blog := service.GetBlogById(blogId)
	blogType := service.GetTypeById(blog.TypeId)
	tags := service.ListTagsByBlogId(blogId)
	blogView := model_view.BlogView{
		Blog: blog,
		Type: blogType,
		Tags: tags,
	}
	c.JSON(code, blogView)
}

func GetBlogsCount(c *gin.Context) {
	count := service.GetBlogsCount()
	c.JSON(http.StatusOK, count)
}

func CreateBlog(c *gin.Context) {
	blogView := model_view.BlogView{}
	if err := c.ShouldBind(&blogView); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "error occurred while parsing blogView, please check the input again")
	}
	service.CreateBlog(&blogView.Blog)
	service.SaveAssociationBetweenBlogAndTags(blogView.Blog.BlogId, blogView.Tags)
	c.JSON(http.StatusCreated, nil)
}

func UpdateBlog(c *gin.Context) {
	blog := model.Blog{}
	if err := c.ShouldBind(&blog); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "error occurred while parsing blog, please check the input again")
	}
	if err := service.UpdateBlog(&blog); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusAccepted, nil)
}

func DeleteBlog(c *gin.Context) {
	blogId := c.Param("blog_id")
	if err := service.DeleteBlog(blogId); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusAccepted, nil)
}

func ListBlogsGroupedByYear(c *gin.Context) {
	blogsGroupedByYear := service.ListBlogsGroupedByYear()
	c.JSON(http.StatusOK, blogsGroupedByYear)
}
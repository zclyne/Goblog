package router

import (
	"goblog/model"
	"goblog/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTags(c *gin.Context) {
	tags := service.ListTags()
	c.JSON(http.StatusOK, tags)
}

func CreateTag(c *gin.Context) {
	t := model.Tag{}
	if err := c.ShouldBind(&t); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "error occurred while parsing tag, please check the input again")
	}
	service.CreateTag(&t)
	c.JSON(http.StatusCreated, nil)
}

func DeleteTag(c *gin.Context) {
	tagId := c.Param("tag_id")
	if err := service.DeleteTag(tagId); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusAccepted, nil)
}

func UpdateTag(c *gin.Context) {
	t := model.Tag{}
	if err := c.ShouldBind(&t); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "error occurred while parsing tag, please check the input again")
	}
	if err := service.UpdateTag(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusAccepted, nil)
}

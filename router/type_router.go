package router

import (
	"goblog/model"
	"goblog/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ListTypes(c *gin.Context) {
	types := service.ListTypes()
	c.JSON(http.StatusOK, types)
}

func GetType(c *gin.Context) {
	typeId := c.Param("type_id")
	blogType := service.GetTypeById(typeId)
	c.JSON(http.StatusOK, blogType)
}

func CreateType(c *gin.Context) {
	t := model.Type{}
	if err := c.ShouldBind(&t); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "error occurred while parsing type, please check the input again")
	}
	service.CreateType(&t)
	c.JSON(http.StatusCreated, nil)
}

func DeleteType(c *gin.Context) {
	typeId := c.Param("type_id")
	if err := service.DeleteType(typeId); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusAccepted, nil)
}

func UpdateType(c *gin.Context) {
	t := model.Type{}
	if err := c.ShouldBind(&t); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, "error occurred while parsing type, please check the input again")
	}
	if err := service.UpdateType(&t); err != nil {
		c.JSON(http.StatusBadRequest, err)
	}
	c.JSON(http.StatusAccepted, nil)
}

package router

import (
	"goblog/model"
	"goblog/service"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserRegister(c *gin.Context) {
	user := model.User{}
	if err := c.ShouldBind(&user); err != nil {
		log.Printf("Error parsing the request body to model.User, err is %s\n", err)
		c.JSON(http.StatusNotAcceptable, map[string]string{
			"msg": "There is some error with the uploaded request body, please check the input again",
		})
	}
	service.Register(&user)
}

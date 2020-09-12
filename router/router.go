package router

import (
	"fmt"
	"goblog/config"
	"goblog/model"
	"goblog/service"
	"log"
	"net/http"
	"time"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type loginForm struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

const identityKey = "id"

func InitRouter() {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	// configure CORS
	r.Use(cors.Default())
	apis := r.Group("/api")

	// create jwt middleware
	authMiddleware, err := jwt.New(&jwt.GinJWTMiddleware{
		Realm:       "test zone",
		Key:         []byte("secret key"),
		Timeout:     2 * time.Hour,
		MaxRefresh:  time.Hour,
		IdentityKey: identityKey,
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			fmt.Printf("PayloadFunc, data: %+v\n", data)
			if v, ok := data.(*model.User); ok {
				return jwt.MapClaims{
					identityKey: v.Username,
					"type":      v.Type,
					"email":     v.Email,
				}
			}
			return jwt.MapClaims{}
		},
		IdentityHandler: func(c *gin.Context) interface{} {
			claims := jwt.ExtractClaims(c)
			fmt.Printf("IdentityHandler, claims: %+v\n", claims)
			return &model.User{
				Username: claims[identityKey].(string),
				// when directly retrieve from claims, claims["type"] is a float64, so we must cast it into int
				// this might be a problem of json unmarshalling
				Type:  int(claims["type"].(float64)),
				Email: claims["email"].(string),
			}
		},
		Authenticator: func(c *gin.Context) (interface{}, error) {
			var loginVals loginForm
			if err := c.ShouldBind(&loginVals); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			username := loginVals.Username
			password := loginVals.Password

			if user := service.Login(username, password); user.UserId != "" {
				fmt.Printf("Authenticator: %+v\n", user)
				return &user, nil
			}

			return nil, jwt.ErrFailedAuthentication
		},
		Authorizator: func(data interface{}, c *gin.Context) bool {
			v, ok := data.(*model.User)
			fmt.Printf("Authorizator: %+v\n", v)
			if ok && v.Type == 1 {
				return true
			}

			return false
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
	})
	if err != nil {
		log.Fatal("JWT Error:" + err.Error())
	}

	// user api
	apis.POST("/user/login", authMiddleware.LoginHandler)
	apis.GET("/user/refresh_token", authMiddleware.RefreshHandler)
	apis.POST("/user/logout", authMiddleware.LogoutHandler)
	apis.POST("/user/register", UserRegister)

	// authMiddlware only takes effect on the routers added after the authMiddlware is set
	// apis.Use(authMiddleware.MiddlewareFunc())

	// blogs api
	apis.GET("/blogs", ListBlogs)
	apis.POST("/blogs", CreateBlog)
	apis.GET("/blogs/:blog_id", GetBlog)
	apis.GET("/blogs-count", GetBlogsCount)
	apis.GET("blogs-archive", ListBlogsGroupedByYear)

	// types api
	apis.GET("/types", ListTypes)
	apis.GET("/types/:type_id", GetType)
	apis.POST("/types", CreateType)
	apis.DELETE("/types/:type_id", DeleteType)
	apis.PUT("/types", UpdateType)

	// tags api
	apis.GET("/tags", ListTags)
	apis.POST("/tags", CreateTag)
	apis.DELETE("/tags/:tag_id", DeleteTag)
	apis.PUT("/tags", UpdateTag)

	// comments api
	apis.GET("/comments", ListCommentsByBlogId)
	apis.POST("/comments", CreateComment)
	apis.DELETE("/comments/:comment_id", DeleteComment)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    "CONTENT_NOT_FOUND",
			"message": "Content not found",
		})
	})

	r.Run(":" + config.Configuration.Port)
}

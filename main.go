package main

import (
	"go-contacts/controllers"
	"github.com/gin-gonic/gin"
	"go-contacts/jwt"
	middlewares "go-contacts/middleware"
	"go-contacts/database"
)

func main() {

	database.Connect()
	r := gin.Default()
	r.Use(middlewares.InjectMongoDB)


	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"body": "ok",
		})
	})

	v1 := r.Group("/api/v1")
	{
		topics := v1.Group("/topics")
		{
			topics.GET("/:_id", controllers.GetTopicByID)
			topics.GET("/", controllers.GetTopics)
			topics.POST("/", controllers.CreateTopic)
			topics.DELETE("/:_id", controllers.DeleteTopic)
		}

		users := v1.Group("/users")
		users.Use(jwt.JWT())
		{
			users.GET("/all",controllers.GetUsers)
			users.GET("/id/:id", controllers.GetUserById)
			users.GET("/username/:name", controllers.GetUserByName)
			users.POST("/new/user", controllers.SaveUser)
			users.DELETE("/:userID", controllers.DeleteUserFormDatabase)
		}
	}


	//endless.ListenAndServe(":"+os.Getenv("PORT"), r)
	r.Run() // listen and serve on 0.0.0.0:8080
}

package routes

import (
	c "gin-mongo-boilerplate/src/user/controllers"

	. "github.com/shipu/artifact"
)

func TodoSetup() {
	v1 := Router.Group("api/v1")
	v1.GET("users", c.TodoIndex())
	v1.POST("users", c.TodoCreate())
	v1.GET("users/:userId", c.TodoShow())
	v1.PUT("users/:userId", c.TodoUpdate())
	v1.DELETE("users/:userId", c.TodoDelete())
}

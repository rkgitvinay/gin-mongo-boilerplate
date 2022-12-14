package routes

import (
	c "gin-mongo-boilerplate/src/user/controllers"

	. "github.com/shipu/artifact"
)

func UserSetup() {
	v1 := Router.Group("api/v1")
	v1.GET("users", c.UserIndex())
	v1.POST("users", c.UserCreate())
	v1.GET("users/:userId", c.UserShow())
	v1.PUT("users/:userId", c.UserUpdate())
	v1.DELETE("users/:userId", c.UserDelete())
}

package controllers

import (
	"gin-mongo-boilerplate/src/user/dto"
	"gin-mongo-boilerplate/src/user/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shipu/artifact"
)

// TodoIndex
// @Summary  all users
// @Schemes
// @Description  All users
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        page    query  string  false  "Page"
// @Param        limit   query  string  false  "Limit"
// @Param        status  query  string  false  "Status"
// @Success      200
// @Failure      401
// @Router       /users [get]
func TodoIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")
		status := c.DefaultQuery("status", "")

		var filter map[string]interface{} = make(map[string]interface{})
		filter["page"] = page
		filter["limit"] = limit
		filter["status"] = status

		users, paginate := services.AllTodo(filter)

		artifact.Res.Code(200).Data(users).Raw(map[string]interface{}{
			"meta": paginate,
		}).Json(c)
	}
}

// TodoCreate
// @Summary  create a user
// @Schemes
// @Description  create a user
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        request  body  dto.CreateTodoRequest  true  "Create Todo Request"
// @Success      200
// @Failure      401
// @Router       /users [post]
func TodoCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createTodo dto.CreateTodoRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createTodo); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user := services.CreateATodo(createTodo)

		artifact.Res.Code(http.StatusCreated).Message("success").Data(user).Json(c)
	}
}

// TodoShow
// @Summary  user details
// @Schemes
// @Description  Todo Details
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        userId  path  string  true  "Todo ID"
// @Success      200
// @Failure      401
// @Router       /users/{userId} [get]
func TodoShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		userId := c.Param("userId")

		user := services.ATodo(userId)

		artifact.Res.Code(http.StatusOK).Message("success").Data(user).Json(c)
	}
}

// TodoUpdate
// @Summary  update a user
// @Schemes
// @Description  update a user
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        userId   path  string                 true  "Todo ID"
// @Param        request  body  dto.UpdateTodoRequest  true  "Update Todo Request"
// @Success      200
// @Failure      401
// @Router       /users/{userId} [put]
func TodoUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateTodo dto.UpdateTodoRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")

		if err := c.ShouldBind(&updateTodo); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user, err := services.UpdateATodo(userId, updateTodo)

		if err != nil {
			artifact.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(user).Json(c)
	}
}

// TodoDelete
// @Summary  delete a user
// @Schemes
// @Description  delete a user
// @Tags         Todo
// @Accept       json
// @Produce      json
// @Param        userId  path  string  true  "Todo ID"
// @Success      200
// @Failure      422
// @Router       /users/{userId} [delete]
func TodoDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")
		err := services.DeleteATodo(userId)

		if !err {
			artifact.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}

package controllers

import (
	"gin-mongo-boilerplate/src/user/dto"
	"gin-mongo-boilerplate/src/user/services"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shipu/artifact"
)

// UserIndex
// @Summary  all users
// @Schemes
// @Description  All users
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        page    query  string  false  "Page"
// @Param        limit   query  string  false  "Limit"
// @Param        status  query  string  false  "Status"
// @Success      200
// @Failure      401
// @Router       /users [get]
func UserIndex() gin.HandlerFunc {
	return func(c *gin.Context) {
		page := c.DefaultQuery("page", "1")
		limit := c.DefaultQuery("limit", "10")
		status := c.DefaultQuery("status", "")

		var filter map[string]interface{} = make(map[string]interface{})
		filter["page"] = page
		filter["limit"] = limit
		filter["status"] = status

		users, paginate := services.AllUser(filter)

		artifact.Res.Code(200).Data(users).Raw(map[string]interface{}{
			"meta": paginate,
		}).Json(c)
	}
}

// UserCreate
// @Summary  create a user
// @Schemes
// @Description  create a user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        request  body  dto.CreateUserRequest  true  "Create User Request"
// @Success      200
// @Failure      401
// @Router       /users [post]
func UserCreate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var createUser dto.CreateUserRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		if err := c.ShouldBind(&createUser); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message("Bad Request").Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user := services.CreateAUser(createUser)

		artifact.Res.Code(http.StatusCreated).Message("success").Data(user).Json(c)
	}
}

// UserShow
// @Summary  user details
// @Schemes
// @Description  User Details
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId  path  string  true  "User ID"
// @Success      200
// @Failure      401
// @Router       /users/{userId} [get]
func UserShow() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusNotFound).Message(http.StatusText(http.StatusNotFound)).Json(c)
			}
		}()

		userId := c.Param("userId")

		user := services.AUser(userId)

		artifact.Res.Code(http.StatusOK).Message("success").Data(user).Json(c)
	}
}

// UserUpdate
// @Summary  update a user
// @Schemes
// @Description  update a user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId   path  string                 true  "User ID"
// @Param        request  body  dto.UpdateUserRequest  true  "Update User Request"
// @Success      200
// @Failure      401
// @Router       /users/{userId} [put]
func UserUpdate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var updateUser dto.UpdateUserRequest

		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message(http.StatusText(http.StatusUnprocessableEntity)).Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")

		if err := c.ShouldBind(&updateUser); err != nil {
			artifact.Res.Code(http.StatusBadRequest).Message(http.StatusText(http.StatusBadRequest)).Data(err.Error()).AbortWithStatusJSON(c)
			return
		}

		user, err := services.UpdateAUser(userId, updateUser)

		if err != nil {
			artifact.Res.Code(http.StatusInternalServerError).Message(http.StatusText(http.StatusInternalServerError)).Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Updated !!!").Data(user).Json(c)
	}
}

// UserDelete
// @Summary  delete a user
// @Schemes
// @Description  delete a user
// @Tags         User
// @Accept       json
// @Produce      json
// @Param        userId  path  string  true  "User ID"
// @Success      200
// @Failure      422
// @Router       /users/{userId} [delete]
func UserDelete() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				artifact.Res.Code(http.StatusUnprocessableEntity).Message("error").Data(err).Json(c)
			}
		}()

		userId := c.Param("userId")
		err := services.DeleteAUser(userId)

		if !err {
			artifact.Res.Code(http.StatusInternalServerError).Message("something wrong").Json(c)
			return
		}

		artifact.Res.Code(http.StatusOK).Message("Successfully Delete !!!").Json(c)
	}
}

package user

import (
	"errors"
	"net/http"
	"strconv"

	"go_backend/app/handler"
	userRepo "go_backend/app/repo/mysql/user"
	"go_backend/domain"
	"go_backend/model"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type userHandler struct {
	UserService domain.IUserService
}

func NewUserHandler(userService domain.IUserService) domain.IUserHandler {
	return &userHandler{
		UserService: userService,
	}
}

func (h *userHandler) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &handler.ErrorResponse{
				Code: handler.ErrRequestInvalid,
				Msg:  "id must be an integer",
			})
			return
		}

		user, err := h.UserService.GetUserByID(c.Request.Context(), uint(userID))
		if err != nil {
			if errors.Is(err, userRepo.ErrUserNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, &handler.ErrorResponse{
					Code: handler.ErrNotFound,
					Msg:  userRepo.ErrUserNotFound.Error(),
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, &handler.ErrorResponse{
				Code: handler.ErrInternalServer,
				Msg:  handler.ErrInternalServerMsg,
			})
			return
		}

		c.JSON(http.StatusOK, &handler.Response{
			Data: user,
		})
	}
}

func (h *userHandler) GetUserList() gin.HandlerFunc {
	return func(c *gin.Context) {
		var req *getUserListReq
		if err := c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &handler.ErrorResponse{
				Code: handler.ErrRequestInvalid,
				Msg:  err.Error(),
			})
			return
		}

		userListResp, err := h.UserService.GetUserList(c.Request.Context(), req.Page, req.Limit)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, &handler.ErrorResponse{
				Code: handler.ErrInternalServer,
				Msg:  handler.ErrInternalServerMsg,
			})
			return
		}

		c.JSON(http.StatusOK, &handler.Response{
			Data: userListResp,
		})
	}
}

type getUserListReq struct {
	Page  int `form:"page" binding:"required,min=1"`
	Limit int `form:"limit" binding:"required,min=1"`
}

func (h *userHandler) UpdateUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &handler.ErrorResponse{
				Code: handler.ErrRequestInvalid,
				Msg:  "id must be an integer",
			})
			return
		}

		// make sure user can only update his/her own info
		if userID != uint64(c.GetInt("userID")) {
			c.AbortWithStatusJSON(http.StatusForbidden, &handler.ErrorResponse{
				Code: handler.ErrForbidden,
				Msg:  "Not allowed to update other user's info",
			})
			return
		}

		var req *updateUserReq
		if err = c.ShouldBindJSON(&req); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &handler.ErrorResponse{
				Code: handler.ErrRequestInvalid,
				Msg:  err.Error(),
			})
			return
		}

		err = h.UserService.UpdateUser(c.Request.Context(), &model.User{
			Model: gorm.Model{
				ID: uint(userID),
			},
			Mobile:   req.Mobile,
			Name:     req.Name,
			Age:      req.Age,
			Password: req.Password,
		})
		if err != nil {
			if errors.Is(err, userRepo.ErrUserNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, &handler.ErrorResponse{
					Code: handler.ErrNotFound,
					Msg:  userRepo.ErrUserNotFound.Error(),
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, &handler.ErrorResponse{
				Code: handler.ErrInternalServer,
				Msg:  handler.ErrInternalServerMsg,
			})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

type updateUserReq struct {
	Email    string `form:"email,omitempty" binding:"omitempty,email,max=50"`
	Mobile   string `form:"mobile,omitempty" binding:"omitempty,max=11"`
	Name     string `form:"name,omitempty" binding:"omitempty,max=20"`
	Age      int    `form:"age,omitempty" binding:"omitempty,min=1,max=150"`
	Password string `form:"password,omitempty" binding:"omitempty,min=8,max=20,alphanum"`
}

func (h *userHandler) DeleteUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, err := strconv.ParseUint(c.Param("id"), 10, 0)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, &handler.ErrorResponse{
				Code: handler.ErrRequestInvalid,
				Msg:  "id must be an integer",
			})
			return
		}

		// make sure user can only delete his/her own info
		if userID != uint64(c.GetInt("userID")) {
			c.AbortWithStatusJSON(http.StatusForbidden, &handler.ErrorResponse{
				Code: handler.ErrForbidden,
				Msg:  "Not allowed to delete other user",
			})
			return
		}

		err = h.UserService.DeleteUserByID(c.Request.Context(), uint(userID))
		if err != nil {
			if errors.Is(err, userRepo.ErrUserNotFound) {
				c.AbortWithStatusJSON(http.StatusNotFound, &handler.ErrorResponse{
					Code: handler.ErrNotFound,
					Msg:  userRepo.ErrUserNotFound.Error(),
				})
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, &handler.ErrorResponse{
				Code: handler.ErrInternalServer,
				Msg:  handler.ErrInternalServerMsg,
			})
			return
		}

		c.JSON(http.StatusNoContent, nil)
	}
}

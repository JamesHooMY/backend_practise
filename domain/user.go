package domain

//go:generate mockgen -destination ./mock/user.go -source=./user.go -package=mock

import (
	"context"
	"time"

	"go_backend/model"

	"github.com/gin-gonic/gin"
)

type IUserHandler interface {
	GetUserByID() gin.HandlerFunc
	GetUserList() gin.HandlerFunc
	UpdateUserByID() gin.HandlerFunc
	DeleteUserByID() gin.HandlerFunc
}

type IUserService interface {
	GetUserByID(ctx context.Context, id uint) (userResp *UserResp, err error)
	GetUserList(ctx context.Context, page, limit int) (userListResp *UserListResp, err error)
	UpdateUser(ctx context.Context, user *model.User) (err error)
	DeleteUserByID(ctx context.Context, id uint) (err error)
}

type UserResp struct {
	ID        uint      `json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	DeletedAt time.Time `json:"deletedAt"`
	Email     string    `json:"email"`
	Mobile    string    `json:"mobile"`
	Name      string    `json:"name"`
	Age       int       `json:"age"`
}

type UserListResp struct {
	UserList []*UserResp `json:"userList"`
	Total    int64       `json:"total"`
	Page     int         `json:"page"`
	Limit    int         `json:"limit"`
}

type IUserQueryRepo interface {
	GetUserByEmail(ctx context.Context, email string) (user *model.User, err error)
	GetUserByID(ctx context.Context, id uint) (user *model.User, err error)
	GetUserList(ctx context.Context, page, limit int) (userList []*model.User, total int64, err error)
}

type IUserCommandRepo interface {
	CreateUser(ctx context.Context, user *model.User) (err error)
	UpdateUser(ctx context.Context, user *model.User) (err error)
	DeleteUser(ctx context.Context, id uint) (err error)
}

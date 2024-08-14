package app

import (
	userHdl "go_backend/app/handler/user"
	userRepo "go_backend/app/repo/mysql/user"
	userSrv "go_backend/app/service/user"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRouter(router *gin.Engine, db *gorm.DB) *gin.Engine {
	userHandler := userHdl.NewUserHandler(userSrv.NewUserService(
		userRepo.NewUserQueryRepo(db),
		userRepo.NewUserCommandRepo(db),
	))

	user := router.Group("/user")
	user.GET("/:id", userHandler.GetUserByID())
	user.POST("/userList", userHandler.GetUserList())
	user.PUT("/:id", userHandler.UpdateUserByID())
	user.DELETE("/:id", userHandler.DeleteUserByID())

	return router
}

package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/avito-internship-2023/internal/handler"
	"github.com/linqcod/avito-internship-2023/internal/repository"
	"github.com/linqcod/avito-internship-2023/internal/service"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func InitRouter(ctx context.Context, logger *zap.SugaredLogger, db *sql.DB) *gin.Engine {
	router := gin.Default()
	// init swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// init services, repos, handlers
	userRepo := repository.NewUserRepository(ctx, db)
	segmentRepo := repository.NewSegmentRepository(ctx, db)

	userService := service.NewUserService(userRepo)
	segmentService := service.NewSegmentService(segmentRepo)

	userHandler := handler.NewUserHandler(logger, userService)
	segmentHandler := handler.NewSegmentHandler(logger, segmentService)

	api := router.Group("/api/v1")
	{
		users := api.Group("/users")
		{
			users.POST("", userHandler.CreateUser)
			users.GET("/:id", userHandler.GetUserById)
			users.GET("", userHandler.GetAllUsers)
			users.POST("/:id/changeSegments", userHandler.ChangeUserSegments)
			users.GET("/:id/active", userHandler.GetUserActiveSegments)
		}
		segments := api.Group("/segments")
		{
			segments.POST("", segmentHandler.CreateSegment)
			segments.DELETE("/:slug", segmentHandler.DeleteSegment)
		}
	}

	return router
}

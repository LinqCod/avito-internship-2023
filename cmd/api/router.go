package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/avito-internship-2023/internal/handler"
	"github.com/linqcod/avito-internship-2023/internal/repository"
	"github.com/linqcod/avito-internship-2023/internal/service"
	"go.uber.org/zap"
)

func InitRouter(ctx context.Context, logger *zap.SugaredLogger, db *sql.DB) *gin.Engine {
	router := gin.Default()
	// TODO: init swagger, router logger

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
			users.Handle("POST", "", userHandler.CreateUser)
			users.Handle("GET", "/:id", userHandler.GetUserById)
			users.Handle("GET", "", userHandler.GetAllUsers)
			users.Handle("POST", "/:id/changeSegments", userHandler.ChangeUserSegments)
			users.Handle("GET", "/:id/active", userHandler.GetUserActiveSegments)
		}
		segments := api.Group("/segments")
		{
			segments.Handle("POST", "", segmentHandler.CreateSegment)
			segments.Handle("DELETE", "/:slug", segmentHandler.DeleteSegment)
		}
	}

	return router
}

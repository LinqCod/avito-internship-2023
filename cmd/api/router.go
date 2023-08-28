package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func InitRouter(ctx context.Context, db *sql.DB) *gin.Engine {
	router := gin.Default()

	// TODO: init services, repos, handlers

	router.Group("/segments")
	{
		// TODO: endpoints
	}

	return router
}

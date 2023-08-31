package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/avito-internship-2023/internal/common/errorTypes"
	"github.com/linqcod/avito-internship-2023/internal/handler/dto"
	"go.uber.org/zap"
	"net/http"
	"path/filepath"
	"strconv"
)

type HistoryService interface {
	GetUserSegmentHistory(userId int64, month int64, year int64) error
}

type HistoryHandler struct {
	logger  *zap.SugaredLogger
	service HistoryService
}

func NewHistoryHandler(logger *zap.SugaredLogger, service HistoryService) *HistoryHandler {
	return &HistoryHandler{
		logger:  logger,
		service: service,
	}
}

// GetUserSegmentHistory godoc
//
//	@Summary		get user segment history
//	@Description	get user segment history by month and date
//	@Tags			users
//	@Param			id	path		int					true	"User id"
//	@Param			month	path		int					true	"month to get history from"
//	@Param			year	path		int					true	"year to get history from"
//	@Success		200		"history csv file received successfully"
//	@Failure		400		{object}	dto.ErrorDTO	"error bad request data"
//	@Failure		500		{object}	dto.ErrorDTO	"error while getting history"
//	@Router			/users/{id}/{month}/{year} [get]
func (h HistoryHandler) GetUserSegmentHistory(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	month, err := strconv.Atoi(c.Param("month"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	year, err := strconv.Atoi(c.Param("year"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	err = h.service.GetUserSegmentHistory(int64(userId), int64(month), int64(year))
	if err != nil {
		h.logger.Errorf("error while getting user's history: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	path, err := filepath.Abs("")
	if err != nil {
		h.logger.Errorf("error while creating history path: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrFailedToCreateFile.Error(),
		})
		return
	}

	filename := fmt.Sprintf("%d.csv", userId)

	c.FileAttachment(path+"/"+filename, filename)
}

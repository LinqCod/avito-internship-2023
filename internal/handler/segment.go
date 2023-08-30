package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/avito-internship-2023/internal/common/errorTypes"
	"github.com/linqcod/avito-internship-2023/internal/model"
	"go.uber.org/zap"
	"net/http"
)

type SegmentService interface {
	CreateSegment(segment model.CreateSegmentDTO) (string, error)
	DeleteSegment(slug string) error
}

type SegmentHandler struct {
	logger  *zap.SugaredLogger
	service SegmentService
}

func NewSegmentHandler(logger *zap.SugaredLogger, service SegmentService) *SegmentHandler {
	return &SegmentHandler{
		logger:  logger,
		service: service,
	}
}

func (h SegmentHandler) CreateSegment(c *gin.Context) {
	var segment model.CreateSegmentDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&segment); err != nil {
		h.logger.Errorf("error while unmarshaling segment body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorTypes.ErrJSONUnmarshalling.Error(),
		})
		return
	}

	slug, err := h.service.CreateSegment(segment)
	if err != nil {
		h.logger.Errorf("error while creating segment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorTypes.ErrDBDataInsertion.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"slug": slug,
	})
}

func (h SegmentHandler) DeleteSegment(c *gin.Context) {
	slug := c.Param("slug")

	if err := h.service.DeleteSegment(slug); err != nil {
		h.logger.Errorf("error while deleting segment: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorTypes.ErrDBDataDeletion,
		})
		return
	}

	c.Status(http.StatusNoContent)
}

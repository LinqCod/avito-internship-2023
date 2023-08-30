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

// CreateSegment godoc
//
//	@Summary		create segment
//	@Description	create segment
//	@Tags			segments
//	@Accept			json
//	@Produce		json
//	@Param			segment	body		model.CreateSegmentDTO	true	"Create segment"
//	@Success		201		{object}	model.CreateSegmentResponse	"segment created successfully"
//	@Failure		400		{object}	model.ErrorDTO	"error bad request data"
//	@Failure		500		{object}	model.ErrorDTO	"error while inserting segment to db table"
//	@Router			/segments [post]
func (h SegmentHandler) CreateSegment(c *gin.Context) {
	var segment model.CreateSegmentDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&segment); err != nil {
		h.logger.Errorf("error while unmarshaling segment body: %v", err)
		c.JSON(http.StatusBadRequest, model.ErrorDTO{
			Error: errorTypes.ErrJSONUnmarshalling.Error(),
		})
		return
	}

	slug, err := h.service.CreateSegment(segment)
	if err != nil {
		h.logger.Errorf("error while creating segment: %v", err)
		c.JSON(http.StatusInternalServerError, model.ErrorDTO{
			Error: errorTypes.ErrDBDataInsertion.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, model.CreateSegmentResponse{
		Slug: slug,
	})
}

// DeleteSegment godoc
//
//	@Summary		delete segment
//	@Description	delete segment by id
//	@Tags			segments
//	@Param			slug	path		string					true	"Segment slug"
//	@Success		204	"segment deleted successfully"
//	@Failure		500		{object}	model.ErrorDTO	"error while deleting segment"
//	@Router			/segments/{id} [delete]
func (h SegmentHandler) DeleteSegment(c *gin.Context) {
	slug := c.Param("slug")

	if err := h.service.DeleteSegment(slug); err != nil {
		h.logger.Errorf("error while deleting segment: %v", err)
		c.JSON(http.StatusInternalServerError, model.ErrorDTO{
			Error: errorTypes.ErrDBDataDeletion.Error(),
		})
		return
	}

	c.Status(http.StatusNoContent)
}

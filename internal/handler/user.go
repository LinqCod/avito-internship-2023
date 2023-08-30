package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/avito-internship-2023/internal/common/errorTypes"
	"github.com/linqcod/avito-internship-2023/internal/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserService interface {
	CreateUser(user model.CreateUserDTO) (int64, error)
	GetUserById(id int64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	ChangeUserSegments(dto model.ChangeUserSegmentsDTO, userId int64) error
	GetUserActiveSegments(userId int64) (*model.ActiveUserSegmentsDTO, error)
}

type UserHandler struct {
	logger  *zap.SugaredLogger
	service UserService
}

func NewUserHandler(logger *zap.SugaredLogger, service UserService) *UserHandler {
	return &UserHandler{
		logger:  logger,
		service: service,
	}
}

func (h UserHandler) CreateUser(c *gin.Context) {
	var user model.CreateUserDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		h.logger.Errorf("error while unmarshaling user body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorTypes.ErrJSONUnmarshalling.Error(),
		})
		return
	}

	userId, err := h.service.CreateUser(user)
	if err != nil {
		h.logger.Errorf("error while creating user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorTypes.ErrDBDataInsertion.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user_id": userId,
	})
}

func (h UserHandler) GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	user, err := h.service.GetUserById(int64(userId))
	if err != nil {
		h.logger.Errorf("error while getting user by id: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.logger.Errorf("error while getting users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

func (h UserHandler) ChangeUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	var dto model.ChangeUserSegmentsDTO

	if err = json.NewDecoder(c.Request.Body).Decode(&dto); err != nil {
		h.logger.Errorf("error while unmarshaling user segments body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorTypes.ErrJSONUnmarshalling.Error(),
		})
		return
	}

	if err = h.service.ChangeUserSegments(dto, int64(userId)); err != nil {
		h.logger.Errorf("error while changing user segments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorTypes.ErrDBDataModification.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

func (h UserHandler) GetUserActiveSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"error": errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	activeSegments, err := h.service.GetUserActiveSegments(int64(userId))
	if err != nil {
		h.logger.Errorf("error while getting user active segments: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, activeSegments)
}

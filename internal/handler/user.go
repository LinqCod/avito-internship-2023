package handler

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/avito-internship-2023/internal/common/errorTypes"
	"github.com/linqcod/avito-internship-2023/internal/handler/dto"
	"github.com/linqcod/avito-internship-2023/internal/model"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

type UserService interface {
	CreateUser(user dto.CreateUserDTO) (int64, error)
	GetUserById(id int64) (*model.User, error)
	GetAllUsers() ([]*model.User, error)
	ChangeUserSegments(userSegmDTO dto.ChangeUserSegmentsDTO, userId int64) error
	GetUserActiveSegments(userId int64) (*dto.ActiveUserSegmentsDTO, error)
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

// CreateUser godoc
//
//	@Summary		create user
//	@Description	create user
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			user	body		dto.CreateUserDTO	true	"Create user"
//	@Success		201		{object}	dto.CreateUserResponse	"user created successfully"
//	@Failure		400		{object}	dto.ErrorDTO	"error bad request data"
//	@Failure		500		{object}	dto.ErrorDTO	"error while inserting user to db table"
//	@Router			/users [post]
func (h UserHandler) CreateUser(c *gin.Context) {
	var user dto.CreateUserDTO

	if err := json.NewDecoder(c.Request.Body).Decode(&user); err != nil {
		h.logger.Errorf("error while unmarshaling user body: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrJSONUnmarshalling.Error(),
		})
		return
	}

	userId, err := h.service.CreateUser(user)
	if err != nil {
		h.logger.Errorf("error while creating user: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrDBDataInsertion.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.CreateUserResponse{
		Id: userId,
	})
}

// GetUserById godoc
//
//	@Summary		get user
//	@Description	get user by id
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int					true	"User id"
//	@Success		200		{object}	model.User 	"user received successfully"
//	@Failure		400		{object}	dto.ErrorDTO	"error bad request data"
//	@Failure		500		{object}	dto.ErrorDTO	"error while getting user"
//	@Router			/users/{id} [get]
func (h UserHandler) GetUserById(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	user, err := h.service.GetUserById(int64(userId))
	if err != nil {
		h.logger.Errorf("error while getting user by id: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, user)
}

// GetAllUsers godoc
//
//	@Summary		get users
//	@Description	get all users
//	@Tags			users
//	@Produce		json
//	@Success		200		{array}	model.User 	"all users received successfully"
//	@Failure		500		{object}	dto.ErrorDTO	"error while getting users"
//	@Router			/users [get]
func (h UserHandler) GetAllUsers(c *gin.Context) {
	users, err := h.service.GetAllUsers()
	if err != nil {
		h.logger.Errorf("error while getting users: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, users)
}

// ChangeUserSegments godoc
//
//	@Summary		change user segments
//	@Description	add and remove user segments
//	@Tags			users
//	@Accept			json
//	@Produce		json
//	@Param			segmentsChanges	body		dto.ChangeUserSegmentsDTO	true	"Change segments"
//	@Param			id	path		int					true	"User id"
//	@Success		200		"segments changed successfully"
//	@Failure		400		{object}	dto.ErrorDTO	"error bad request data"
//	@Failure		500		{object}	dto.ErrorDTO	"error while changing segments"
//	@Router			/users/{id}/changeSegments [post]
func (h UserHandler) ChangeUserSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	var userSegmDTO dto.ChangeUserSegmentsDTO

	if err = json.NewDecoder(c.Request.Body).Decode(&userSegmDTO); err != nil {
		h.logger.Errorf("error while unmarshaling user segments body: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrJSONUnmarshalling.Error(),
		})
		return
	}

	h.logger.Infof("%v", userSegmDTO)

	if err = h.service.ChangeUserSegments(userSegmDTO, int64(userId)); err != nil {
		h.logger.Errorf("error while changing user segments: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrDBDataModification.Error(),
		})
		return
	}

	c.Status(http.StatusOK)
}

// GetUserActiveSegments godoc
//
//	@Summary		get active segments
//	@Description	get user active segments
//	@Tags			users
//	@Produce		json
//	@Param			id	path		int					true	"User id"
//	@Success		200		{object}	dto.ActiveUserSegmentsDTO 	"segments received successfully"
//	@Failure		400		{object}	dto.ErrorDTO	"error bad request data"
//	@Failure		500		{object}	dto.ErrorDTO	"error while getting users"
//	@Router			/users/{id}/active [get]
func (h UserHandler) GetUserActiveSegments(c *gin.Context) {
	userId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		h.logger.Errorf("error while converting string to int: %v", err)
		c.JSON(http.StatusBadRequest, dto.ErrorDTO{
			Error: errorTypes.ErrBadRequestData.Error(),
		})
		return
	}

	activeSegments, err := h.service.GetUserActiveSegments(int64(userId))
	if err != nil {
		h.logger.Errorf("error while getting user active segments: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, activeSegments)
}

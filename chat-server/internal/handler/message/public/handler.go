// nolint
package public

import (
	"context"
	"fmt"
	"github.com/ew0s/ewos-to-go-hw/chat-server/internal/domain/entity"
	"github.com/ew0s/ewos-to-go-hw/chat-server/internal/handler"
	"github.com/ew0s/ewos-to-go-hw/chat-server/internal/handler/mapper"
	handlerinternalutils "github.com/ew0s/ewos-to-go-hw/chat-server/internal/pkg/utils/handler"
	"github.com/go-playground/validator/v10"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/sirupsen/logrus"

	"github.com/ew0s/ewos-to-go-hw/chat-server/internal/handler/request"

	handlerutils "github.com/ew0s/ewos-to-go-hw/chat-server/pkg/utils/handler"
	sliceutils "github.com/ew0s/ewos-to-go-hw/chat-server/pkg/utils/slice"
)

type MessageService interface {
	SendPublicMessage(ctx context.Context, msg entity.PublicMessage) (*entity.PublicMessage, error)
	GetAllPublicMessages(ctx context.Context, offset, limit int) []*entity.PublicMessage
}

type UserService interface {
	RegisterUser(ctx context.Context, user entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, id int) (*entity.User, error)
	GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetAllUsers(ctx context.Context, offset, limit int) []*entity.User
	UpdateUser(ctx context.Context, id int, updateModel entity.User) (*entity.User, error)
	DeleteUser(ctx context.Context, id int) (*entity.User, error)
}

type Middleware = func(http.Handler) http.Handler

type Handler struct {
	MessageService MessageService
	UserService    UserService
	Middlewares    []Middleware

	logger    *logrus.Logger
	validator *validator.Validate
}

func New(
	messageService MessageService,
	userService UserService,
	logger *logrus.Logger,
	validator *validator.Validate,
	middlewares ...Middleware,
) *Handler {
	return &Handler{
		MessageService: messageService,
		UserService:    userService,
		Middlewares:    middlewares,
		logger:         logger,
		validator:      validator,
	}
}

func (h *Handler) Routes() *chi.Mux {
	router := chi.NewRouter()

	router.Group(func(r chi.Router) {
		r.Use(h.Middlewares...)

		r.Get("/", h.GetAllPublicMessages)
		r.Post("/", h.SendPublicMessage)
	})

	return router
}

// GetAllPublicMessages godoc
//
//	@Summary		Get all public messages
//	@Description	Get all public messages that were sent to chat
//	@Security		BasicAuth
//	@Security		JWT
//	@Tags			Message
//	@Produce		json
//	@Param			offset	query		int	true	"Offset"
//	@Param			limit	query		int	true	"Limit"
//	@Success		200		{object}	[]response.GetPublicMessageResponse
//	@Failure		401		{string}	Unauthorized
//	@Router			/api/v1/messages/public [get]
func (h *Handler) GetAllPublicMessages(rw http.ResponseWriter, req *http.Request) {
	paginationOpts := handlerinternalutils.GetPaginationOptsFromQuery(req, handler.DefaultOffset, handler.DefaultLimit)

	if err := paginationOpts.Validate(h.validator); err != nil {
		handlerutils.WriteErrResponseAndLog(rw, h.logger, http.StatusBadRequest, "", err.Error())

		return
	}

	messages := h.MessageService.GetAllPublicMessages(req.Context(), paginationOpts.Offset, paginationOpts.Limit)

	render.JSON(rw, req, sliceutils.Map(messages, mapper.MapPublicMessageToResponse))
	rw.WriteHeader(http.StatusOK)
}

// SendPublicMessage godoc
//
//	@Summary		Send public message to chat
//	@Description	Send public message to chat
//	@Security		BasicAuth
//	@Security		JWT
//	@Tags			Message
//	@Accept			json
//	@Produce		json
//	@Param			input	body		request.SendPublicMessageRequest	true	"public message schema"
//	@Success		200		{object}	[]response.GetPublicMessageResponse
//	@Failure		401		{string}	Unauthorized
//	@Failure		500		{string}	internal	error
//	@Router			/api/v1/messages/public [post]
func (h *Handler) SendPublicMessage(rw http.ResponseWriter, req *http.Request) {
	username, err := handlerutils.GetStringHeaderByKey(req, "username")
	if err != nil {
		handlerutils.WriteErrResponseAndLog(rw, h.logger, http.StatusUnauthorized, "", err.Error())
		return
	}

	var pubMsgReq request.SendPublicMessageRequest

	if err = render.DecodeJSON(req.Body, &pubMsgReq); err != nil {
		logMsg := fmt.Sprintf("error occurred validating PublicMessageRequest struct: %s", err)
		respMsg := fmt.Sprintf("invalid message provided: %s", err)

		handlerutils.WriteErrResponseAndLog(rw, h.logger, http.StatusBadRequest, logMsg, respMsg)

		return
	}

	if err = pubMsgReq.Validate(h.validator); err != nil {
		logMsg := fmt.Sprintf("error occurred validating PublicMessageRequest struct: %s", err)
		respMsg := fmt.Sprintf("invalid message provided: %s", err)

		handlerutils.WriteErrResponseAndLog(rw, h.logger, http.StatusBadRequest, logMsg, respMsg)

		return
	}

	message, err := h.MessageService.SendPublicMessage(req.Context(), mapper.MapSendPublicMessageRequestToEntity(pubMsgReq, username))
	if err != nil {
		logMsg := fmt.Sprintf("error occurred saving public message: %s", err)

		handlerutils.WriteErrResponseAndLog(rw, h.logger, http.StatusInternalServerError, logMsg, "")

		return
	}

	render.JSON(rw, req, mapper.MapPublicMessageToResponse(message))
	rw.WriteHeader(http.StatusCreated)
}

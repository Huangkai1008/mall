package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go.uber.org/zap"

	"mall/internal/app/v1/user"
	"mall/internal/app/v1/user/schema"
	"mall/internal/app/v1/user/service"
	resp "mall/internal/pkg/util/response"
)

type Handler struct {
	logger  *zap.Logger
	service *service.Service
}

func NewHandler(logger *zap.Logger, service *service.Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

func (h *Handler) Register(c *gin.Context) {
	// Validate
	var registerSchema schema.RegisterSchema
	if err := c.ShouldBind(&registerSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		resp.BadEntityRequest(c, registerSchema.Validate(errs))
		return
	}

	u := user.User{
		Username: registerSchema.Username,
		Email:    registerSchema.Email,
		Password: registerSchema.Password,
	}
	if u, err := h.service.Create(&u); err != nil {
		resp.BadRequest(c, err.Error())
		return
	} else {
		resp.Created(c, u)
	}
}

var ProviderSet = wire.NewSet(NewHandler)

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"go.uber.org/zap"

	"mall/internal/app/v1/account"
	"mall/internal/app/v1/account/schema"
	"mall/internal/app/v1/account/service"
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

// Create Account with register action.
func (h *Handler) Create(c *gin.Context) {
	var createSchema schema.CreateSchema
	if err := c.ShouldBind(&createSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		resp.BadEntityRequest(c, createSchema.Validate(errs))
		return
	}

	a := account.Account{
		Username: createSchema.Username,
		Email:    createSchema.Email,
		Password: createSchema.Password,
	}
	if a, err := h.service.Create(&a); err != nil {
		resp.BadRequest(c, err.Error())
		return
	} else {
		resp.Created(c, a)
	}
}

var ProviderSet = wire.NewSet(NewHandler)

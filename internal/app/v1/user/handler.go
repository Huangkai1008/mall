package user

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"

	res "mall/internal/pkg/util/response"
)

type Handler struct {
	logger  *zap.Logger
	service *Service
}

func NewHandler(logger *zap.Logger, service *Service) *Handler {
	return &Handler{
		logger:  logger,
		service: service,
	}
}

// Register user register.
func (h *Handler) Register(c *gin.Context) {
	// Validate
	var registerSchema RegisterSchema
	if err := c.ShouldBind(&registerSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		res.BadEntityRequest(c, registerSchema.Validate(errs))
		return
	}

	user := User{
		Username: registerSchema.Username,
		Email:    registerSchema.Email,
		Password: registerSchema.Password,
	}
	if user, err := h.service.Create(&user); err != nil {
		res.BadRequest(c, err.Error())
		return
	} else {
		res.Created(c, user)
	}
}

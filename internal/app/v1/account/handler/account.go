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

type AccountHandler struct {
	logger  *zap.Logger
	service *service.AccountService
}

func NewAccountHandler(logger *zap.Logger, service *service.AccountService) *AccountHandler {
	return &AccountHandler{
		logger:  logger,
		service: service,
	}
}

// Register account.
func (h *AccountHandler) Register(c *gin.Context) {
	var regSchema schema.AccountRegisterSchema
	if err := c.ShouldBind(&regSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		resp.BadEntityRequest(c, regSchema.Validate(errs))
		return
	}

	if a, err := h.service.Create(&account.Account{
		Username: regSchema.Username,
		Email:    regSchema.Email,
		Password: regSchema.Password,
	}); err != nil {
		resp.BadRequest(c, err.Error())
		return
	} else {
		resp.Created(c, a)
	}
}

//// Login account.
//func (h *AccountHandler) Login(c *gin.Context) {
//	var loginSchema schema.AccountLoginSchema
//	if err := c.ShouldBind(&loginSchema); err != nil {
//		errs := err.(validator.ValidationErrors)
//		resp.BadEntityRequest(c, loginSchema.Validate(errs))
//		return
//	}
//
//	a := account.Account{
//		Username: loginSchema.Username,
//		Password: loginSchema.Password,
//	}
//	if a, err := h.service.Login(loginSchema.Username, loginSchema.Password); err != nil {
//		resp.BadRequest(c, err.Error())
//		return
//	} else {
//		resp.Created(c, a)
//	}
//}

var ProviderSet = wire.NewSet(NewAccountHandler)

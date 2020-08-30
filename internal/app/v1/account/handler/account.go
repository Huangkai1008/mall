package handler

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
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
func (h *AccountHandler) Register(c echo.Context) (err error) {
	s := new(schema.AccountRegSchema)
	if err = c.Bind(s); err != nil {
		return err
	}
	if err = c.Validate(s); err != nil {
		return err
	}

	if a, err := h.service.Create(&account.Account{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
	}); err != nil {
		return err
	} else {
		return resp.Created(c, a)
	}
}

//// Login account.
//func (h *AccountHandler) Login(c *gin.Context) {
//	var loginSchema schema.AccountLoginSchema
//	if err := c.ShouldBind(&loginSchema); err != nil {
//		errs := err.(validators.ValidationErrors)
//		resp.BadEntityRequest(c, loginSchema.Validate(errs))
//		return
//	}
//
//	if a, err := h.service.Login(loginSchema.Username, loginSchema.Password); err != nil {
//		resp.BadRequest(c, err.Error())
//		return
//	} else {
//		resp.Created(c, a)
//	}
//}

var ProviderSet = wire.NewSet(NewAccountHandler)

package handler

import (
	"github.com/google/wire"
	"github.com/labstack/echo/v4"
	"go.uber.org/zap"

	kitutil "github.com/Huangkai1008/micro-kit/pkg/util"

	"github.com/Huangkai1008/mall/internal/app/v1/account"
	"github.com/Huangkai1008/mall/internal/app/v1/account/schema"
	"github.com/Huangkai1008/mall/internal/app/v1/account/service"
)

type AccountHandler struct {
	*service.AccountService
	logger *zap.Logger
}

func NewAccountHandler(logger *zap.Logger, service *service.AccountService) *AccountHandler {
	return &AccountHandler{
		AccountService: service,
		logger:         logger,
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

	if a, err := h.Create(&account.Account{
		Username: s.Username,
		Email:    s.Email,
		Password: s.Password,
	}); err != nil {
		return err
	} else {
		return kitutil.Created(c, a)
	}
}

// Login account.
func (h *AccountHandler) Login(c echo.Context) (err error) {
	s := new(schema.AccountLoginSchema)
	if err = c.Bind(s); err != nil {
		return err
	}
	if err = c.Validate(s); err != nil {
		return err
	}

	if a, err := h.AccountService.Login(s.Username, s.Password); err != nil {
		return err
	} else {
		return kitutil.Created(c, a)
	}
}

var ProviderSet = wire.NewSet(NewAccountHandler)

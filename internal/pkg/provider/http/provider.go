package http

import (
	"github.com/google/wire"

	"github.com/Huangkai1008/micro-kit/pkg/transport/http"
)

var ProviderSet = wire.NewSet(NewHTTPServer, http.NewRouter, NewValidator)

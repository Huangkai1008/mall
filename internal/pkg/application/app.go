package application

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"go.uber.org/zap"

	"mall/internal/pkg/config"
	"mall/internal/pkg/constant"
	"mall/internal/pkg/logging"
)

// Application is the mall service instance.
// It provides HTTP service now.
type Application struct {
	config     *config.Config
	logger     *zap.Logger
	router     *gin.Engine
	httpServer *http.Server
}

// New returns a new Application.
func New() (*Application, error) {
	// Config
	conf, err := config.New()
	if err != nil {
		return nil, errors.Wrap(err, constant.LoadConfigError)
	}

	// Logging
	logger, err := logging.New(&logging.Options{Config: conf})
	if err != nil {
		return nil, errors.Wrap(err, constant.LogConfigError)
	}

	// Router
	gin.SetMode(conf.RunMode)
	r := gin.New()

	// Application
	application := &Application{
		config: conf,
		logger: logger.With(zap.String("type", "Application")),
		router: r,
	}
	return application, nil
}

// Start Application
func (app *Application) Start() error {
	app.httpServer = &http.Server{
		Addr:           app.config.Addr(),
		Handler:        app.router,
		ReadTimeout:    app.config.ReadTimeout * time.Second,
		WriteTimeout:   app.config.WriteTimeout * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	app.logger.Info("HTTP server starting ...", zap.String("addr", app.config.Addr()))

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.logger.Fatal("Start HTTP server error", zap.Error(err))
		}
		return
	}()
	return nil
}

// Stop Application
func (app *Application) Stop() error {
	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := app.httpServer.Shutdown(ctx); err != nil {
		return errors.Wrap(err, "Shutdown HTTP server error")
	}
	app.logger.Info("Server exiting ...")
	return nil
}

// AwaitSignal await the signal to stop the server
func (app *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		app.logger.Info("Receive a signal", zap.String("signal", s.String()))
		if app.httpServer != nil {
			if err := app.Stop(); err != nil {
				app.logger.Warn("Stop HTTP server error", zap.Error(err))
			}
		}
		os.Exit(0)
	}
}

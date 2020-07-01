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
	"gorm.io/gorm"

	"mall/internal/app/v1/index"
	"mall/internal/app/v1/user"
	"mall/internal/pkg/config"
	"mall/internal/pkg/constant"
	gormApi "mall/internal/pkg/database/gorm"
	"mall/internal/pkg/logging"
	"mall/internal/pkg/middleware"
	repo "mall/internal/pkg/repository"
)

// Application is the mall service instance.
// It provides HTTP service now.
type Application struct {
	config     *config.Config
	logger     *zap.Logger
	db         *gorm.DB
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

	// Database
	db, err := gormApi.New(&gormApi.Options{Config: conf})
	if err != nil {
		return nil, errors.Wrap(err, constant.DatabaseConfigError)
	}

	// Router
	gin.SetMode(conf.RunMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.LoggerMiddleware(logger))

	// Application
	application := &Application{
		config: conf,
		logger: logger.With(zap.String("type", "Application")),
		db:     db,
		router: r,
	}

	// Add apps
	if application.configureApps() != nil {
		return nil, errors.Wrap(err, constant.AppConfigError)
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

func (app *Application) configureApps() error {
	// Repository
	user.NewRepository(app.logger, app.db)

	// Service
	userService := user.NewService(app.logger, repo.GormRepository{
		Logger: app.logger.With(zap.String("type", "UserRepository")),
		Db:     app.db,
	})

	// Handler
	indexHandler := index.NewHandler()
	userHandler := user.NewHandler(app.logger, *userService)

	// Init router
	indexRouter := index.NewRouter(indexHandler)
	userRouter := user.NewRouter(userHandler)

	// API router group
	apiGroup := app.router.Group("/api")
	v1Group := apiGroup.Group("/v1")
	{
		indexRouter(v1Group)
		userRouter(v1Group)
	}
	return nil
}

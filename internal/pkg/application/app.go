package application

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/minio/minio-go/v7"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"

	"mall/internal/pkg/constant"
	"mall/internal/pkg/transport/http"
	"mall/pkg/app"
)

type AppName string

// Application is the application service.
// It contains all the services and configurations.
type Application struct {
	app.App
	logger     *zap.Logger
	router     *gin.Engine
	httpServer *http.Server
	grpcServer *grpc.Server
	minioCli   *minio.Client
	ctx        context.Context
	cancel     context.CancelFunc
}

// Option is the option of application.
type Option func(app *Application) error

// New returns a new Application.
func New(name string, logger *zap.Logger, options ...Option) (*Application, error) {
	a := &Application{
		App: app.App{
			Name: name,
		},
		logger: logger.With(zap.String("type", "Application")),
	}

	for _, option := range options {
		if err := option(a); err != nil {
			return nil, err
		}
	}

	return a, nil
}

// WithHttpServer sets the http server.
func WithHttpServer(s *http.Server) Option {
	return func(a *Application) error {
		a.httpServer = s
		return nil
	}
}

// WithGrpcServer sets the grpc server.
func WithGrpcServer(s *grpc.Server) Option {
	return func(a *Application) error {
		a.grpcServer = s
		return nil
	}
}

// WithMinioCli sets the minio client.
func WithMinioCli(c *minio.Client) Option {
	return func(a *Application) error {
		a.minioCli = c
		return nil
	}
}

//// New returns a new Application.
//func New() (*Application, error) {
//	// Config
//	conf, err := config.New()
//	if err != nil {
//		return nil, errors.Wrap(err, constant.LoadConfigError)
//	}
//
//	// Logging
//	logger, err := logging.New(&logging.Options{Config: conf})
//	if err != nil {
//		return nil, errors.Wrap(err, constant.LogConfigError)
//	}
//
//	// Database
//	db, err := gormApi.New(&gormApi.Options{Config: conf})
//	if err != nil {
//		return nil, errors.Wrap(err, constant.DatabaseConfigError)
//	}
//
//	// Register Translation
//	if err = validator.RegisterTranslation(conf.Locale); err != nil {
//		return nil, errors.Wrap(err, constant.TransRegisterError)
//	}
//
//	// Router
//	gin.SetMode(conf.RunMode)
//	r := gin.New()
//	r.Use(gin.Recovery())
//	r.Use(middleware.LoggerMiddleware(logger))
//
//	// Minio Client
//	minioCli, err := minioApi.New(&minioApi.Options{Config: conf})
//	if err != nil {
//		return nil, errors.Wrap(err, constant.MinioConfigError)
//	}
//
//	// Application
//	application := &Application{
//		config:   conf,
//		logger:   logger.With(zap.String("type", "Application")),
//		db:       db,
//		router:   r,
//		minioCli: minioCli,
//	}
//
//	// Add Apps
//	if application.configureApps() != nil {
//		return nil, errors.Wrap(err, constant.AppConfigError)
//	}
//
//	return application, nil
//}

// Start Application.
func (a *Application) Start() error {
	if err := a.httpServer.Start(); err != nil {
		return errors.Wrap(err, constant.HTTPServerStartError)
	}
	return nil
}

// Stop Application.
func (a *Application) Stop() error {
	if err := a.httpServer.Stop(); err != nil {
		return errors.Wrap(err, constant.HTTPServerStopError)
	}
	a.logger.Info("Server exiting ...")
	return nil
}

// AwaitSignal await the signal to stop the server
func (a *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)

	s := <-c
	a.logger.Info("Receive a signal", zap.String("signal", s.String()))
	if a.httpServer != nil {
		if err := a.Stop(); err != nil {
			a.logger.Warn("Stop HTTP server error", zap.Error(err))
		}
	}
	os.Exit(0)
}

//func (app *Application) configureApps() error {
//	// Repository
//	userRepository := user.NewRepository(app.logger, app.db)
//
//	// Service
//	storageService := storage.NewService(app.config, app.logger, app.minioCli)
//	userService := user.NewService(app.logger, userRepository)
//
//	// Handler
//	indexHandler := index.NewHandler()
//	storageHandler := storage.NewHandler(app.logger, storageService)
//	userHandler := user.NewHandler(app.logger, userService)
//
//	// Init router
//	indexRouter := index.NewRouter(indexHandler)
//	storageRouter := storage.NewRouter(storageHandler)
//	userRouter := user.NewRouter(userHandler)
//
//	// API router group
//	apiGroup := app.router.Group("/api")
//	v1Group := apiGroup.Group("/v1")
//	{
//		indexRouter(v1Group)
//		storageRouter(v1Group)
//		userRouter(v1Group)
//	}
//	return nil
//}

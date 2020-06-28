package application

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"

	"mall/internal/pkg/config"
	"mall/internal/pkg/constant"
)

// Application is the mall service instance.
// It provides HTTP service now.
type Application struct {
	config     *config.Config
	router     *gin.Engine
	httpServer *http.Server
}

// New returns a new Application.
func New() (*Application, error) {
	// Config
	conf, err := config.New()
	if err != nil {
		return nil, errors.Wrap(err, constant.ConfigError)
	}

	// Router
	gin.SetMode(conf.RunMode)
	r := gin.New()

	// Application
	application := &Application{
		config: conf,
		router: r,
	}
	return application, nil
}

// Start Application
func (app *Application) Start() error {
	app.httpServer = &http.Server{
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatal("Start HTTP server error")
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
	log.Println("Server exiting ...")
	return nil
}

// AwaitSignal await the signal to stop the server
func (app *Application) AwaitSignal() {
	c := make(chan os.Signal, 1)
	signal.Reset(syscall.SIGTERM, syscall.SIGINT)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGINT)
	select {
	case s := <-c:
		log.Printf("Receive a signal, signal is %s", s.String())
		if app.httpServer != nil {
			if err := app.Stop(); err != nil {
				log.Println("Stop HTTP server error")
			}
		}
		os.Exit(0)
	}
}

package app

import "github.com/spf13/cobra"

// App is the main structure of a cli application.
type App struct {
	// The cli app's name.
	//
	// Required: true
	Name string

	// The app's description.
	Description string

	// The app's command instance.
	//
	// Required: true
	Command *cobra.Command
}

// NewApp returns a new cli app.
func NewApp(name string, opts ...Option) *App {
	app := &App{
		Name: name,
	}
	for _, o := range opts {
		o(app)
	}

	app.buildCommand()
	return app
}

type Option func(*App)

func WithDescription(description string) Option {
	return func(app *App) {
		app.Description = description
	}
}

func (a *App) buildCommand() {
	a.Command = &cobra.Command{
		Use:   a.Name,
		Short: a.Name,
		Long:  a.Description,
	}
}

package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
)

func New(logger Logger, components ...LifecycleComponent) (*Application, error) {
	return &Application{
		logger:     logger,
		components: components,
	}, nil
}

type Application struct {
	logger     Logger
	components []LifecycleComponent
}

func (a *Application) Run(ctx context.Context) error {
	errChan := make(chan error, len(a.components))

	for _, component := range a.components {
		go func(c LifecycleComponent) {
			a.logger.Info("Starting component:", c.string)
			err := c.Start(ctx)
			if err != nil {
				a.logger.Error("Error starting component %s: %v", c.string, err)
				errChan <- err
			}
		}(component)
	}

	select {
	case <-a.waiting():
		a.logger.Info("App is stopping...")
	case err := <-errChan:
		a.logger.Error("Error occurred, stopping app: %v", err)
		return err
	}

	for _, component := range a.components {
		err := component.Stop(ctx)
		if err != nil {
			a.logger.Error("Error stopping component %s: %v", component.string, err)
			return err
		}
	}

	return nil
}

func (a *Application) waiting() <-chan os.Signal {
	a.logger.Info("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	return wait
}

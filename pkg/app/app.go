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
	for _, component := range a.components {
		err := component.Start(ctx)
		if err != nil {
			return err
		}
	}

	a.waiting()

	for _, component := range a.components {
		err := component.Stop(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *Application) waiting() {
	a.logger.Info("App started!")

	wait := make(chan os.Signal, 1)
	signal.Notify(wait, os.Interrupt, syscall.SIGTERM)

	<-wait

	a.logger.Info("App is stopping...")
}

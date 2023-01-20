package application

import (
	"context"
	"fmt"

	"go.uber.org/zap"

	cfg "github.com/Nikolay-Yakushev/mango/pkg/config"
	httpapp "github.com/Nikolay-Yakushev/mango/internal/adapters/http"
	ports "github.com/Nikolay-Yakushev/mango/internal/ports/driver"
)

type App struct {
	log          *zap.Logger
	description string
	cfg         *cfg.Config
	closeables  []ports.Closeable
}

func (a *App) GetDescription() string {
	return a.description
}

func New(logger *zap.Logger, cfg *cfg.Config) (*App, error) {
	var closeables []ports.Closeable

	a := &App{
		log: logger, 
		description: "Mango component",
		cfg: cfg,
		closeables: closeables,
	}
	return a, nil
}

func (a *App) Start() error {
	webapp, err := httpapp.New(a.cfg, a.log)
	if err != nil {
		a.log.Sugar().Errorw("Failed to start webapp", "reason", err)
		err = fmt.Errorf("Failed to start webapp. Reason %w", err)
		return err
	}
	a.closeables = append(a.closeables, webapp)
	webapp.Start()
	return nil
}

func (a *App) Stop(ctx context.Context) error {
	defer a.log.Sync()

	for _, entity :=range a.closeables{
		errCh := make(chan error, 1)

		go func(){
			err := entity.Stop(ctx)
			errCh <-err
		}()

		select {

			case <-ctx.Done():
				return ctx.Err()

			case err := <-errCh:
				if err != nil{
					// TODO Does this threadsafe? Would i get correct `entity`?
					a.log.Sugar().Error(
						"Failed to stop component %s", entity.GetDescription())
					continue
				}
				a.log.Sugar().Info("Successefully stopped component=%s", entity.GetDescription())
				
		}
		
	}
	return nil
}

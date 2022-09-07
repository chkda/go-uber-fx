package bundlefx

import (
	"context"
	"net/http"

	"github.com/chkda/go-uber-fx-tutorial/configfx"
	"github.com/chkda/go-uber-fx-tutorial/httpfx"
	"github.com/chkda/go-uber-fx-tutorial/loggerfx"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func registerHooks(lifecycle fx.Lifecycle, logger *zap.SugaredLogger, config *configfx.ApplicationConfig, mux *http.ServeMux) {
	lifecycle.Append(
		fx.Hook{
			OnStart: func(context.Context) error {
				logger.Info("Listening on", config.Address)
				go http.ListenAndServe(config.Address, mux)
				return nil
			},
			OnStop: func(context.Context) error {
				return logger.Sync()
			},
		},
	)
}

var Module = fx.Options(
	configfx.Module,
	loggerfx.Module,
	httpfx.Module,
	fx.Invoke(registerHooks),
)

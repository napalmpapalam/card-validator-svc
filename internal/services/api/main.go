package api

import (
	"context"
	"time"

	"github.com/go-chi/chi"
	"github.com/napalmpapalam/card-validator-svc/internal/config"
	"github.com/napalmpapalam/card-validator-svc/internal/services/api/handlers"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/logan/v3"
)

func Run(ctx context.Context, cfg config.Config) {
	r := chi.NewRouter()

	const slowRequestDurationThreshold = time.Second
	ape.DefaultMiddlewares(r, cfg.Log(), slowRequestDurationThreshold)

	r.Use(
		ape.CtxMiddleware(
			handlers.CtxLog(cfg.Log()),
		),
	)

	r.Route("/v1", func(r chi.Router) {
		r.Route("/cards", func(r chi.Router) {
			r.Post("/", handlers.ValidateCard)
		})
	})

	cfg.Log().WithFields(logan.F{
		"service": "api",
		"addr":    cfg.Listener().Addr(),
	}).Info("starting api")

	ape.Serve(ctx, r, cfg, ape.ServeOpts{})
}

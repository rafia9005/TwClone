package workers

import (
	"context"

	"TwClone/internal/config"
	"TwClone/internal/server"
)

func runHttpWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}

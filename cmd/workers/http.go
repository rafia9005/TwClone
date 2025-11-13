package workers

import (
	"context"

	"TWclone/internal/config"
	"TWclone/internal/server"
)

func runHttpWorker(cfg *config.Config, ctx context.Context) {
	srv := server.NewHttpServer(cfg)
	go srv.Start()

	<-ctx.Done()
	srv.Shutdown()
}

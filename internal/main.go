package internal

import (
	"context"
	"errors"
	"log/slog"
	"time"

	"github.com/archey347/dynamic-dns/dynamic-dns/internal/http"
	"github.com/coreos/go-systemd/daemon"
	"github.com/go-chi/chi"
	"golang.org/x/sync/errgroup"
)

func Start(config *Config, log *slog.Logger) error {
	s := http.NewServer(&config.Http, RegisterRoutes, log)

	g, ctx := errgroup.WithContext(context.Background())

	g.Go(func() error {
		l := log.With("component", "watchdog")
		l.Info("Starting")
		watchdog(ctx)

		return nil
	})

	g.Go(func() error {
		l := log.With("component", "http-server")
		l.Info("Starting")

		err := s.Start()
		if err != nil {
			l.Error("Failed to start", "error", err.Error())
		}

		return errors.New("Failed to start http server component")
	})

	return g.Wait()
}

func watchdog(ctx context.Context) {
	for {

		select {
		case <-ctx.Done():
			return
		default:
			daemon.SdNotify(false, daemon.SdNotifyWatchdog)
			time.Sleep(1 * time.Second)
		}
	}
}

func RegisterRoutes(r *chi.Mux) {

}

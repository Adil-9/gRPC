package main

import (
	"grpc/internal/app"
	"grpc/internal/config"
	"log/slog"
	"os"
	"os/signal"
	"syscall"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

// for sso start up
func main() {
	cfg := config.MustLoad()
	// fmt.Println(cfg)

	log := setupLogger(cfg.Env)

	log.Info("starting application", slog.Any("cfg", cfg))

	applicaiton := app.New(log, cfg.GRPC.Port, cfg.Storage_path, cfg.TokenTTL)

	go applicaiton.GRPCSrv.MustRun()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	sign := <-stop

	log.Info("stopping application", slog.String("signal", sign.String()))

	applicaiton.GRPCSrv.Stop()

	log.Info("Application stopped")
	//	todo init object of the config
	//	todo init logger
	// 	todo init the applicaiton (the intry point)
	//	todo start grpc server of the app
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}
	return log
}

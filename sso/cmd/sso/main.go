package main

import (
	"fmt"
	"grpc/internal/config"
)

// for sso start up
func main() {
	cfg := config.MustLoad()
	fmt.Println(cfg)
	//	todo init object of the config
	//	todo init logger
	// 	todo init the applicaiton (the intry point)
	//	todo start grpc server of the app
}

func setupLogger(enc string) *slog.Logger {
}

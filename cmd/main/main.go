package main

import (
	"auth_service/internal/config"
	"auth_service/internal/repositiry/storage"
	"auth_service/internal/transport/grpc"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
)

func main() {
	cfg := config.New()
	fmt.Println(cfg)

	stor, err := storage.New(cfg)
	if err != nil {
		panic(err)
	}

	grpcServer := grpc.New(cfg.ServerCfg, stor)
	go grpcServer.MustStart()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	sign := <-ch

	slog.Info("Got signal: ", sign)
	grpcServer.GracefulStop()

}

type A struct {
	Id    int    `db:"id"`
	Email string `db:"email"`
	Pass  string `db:"password"`
}

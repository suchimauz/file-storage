package app

import (
	"context"
	"errors"
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/suchimauz/file-storage/internal/config"
	delivery "github.com/suchimauz/file-storage/internal/delivery/http"
	"github.com/suchimauz/file-storage/internal/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func Run(configPath string) {
	cfg, err := config.Init(configPath)
	if err != nil {
		return
	}

	spew.Dump(cfg)

	handlers := delivery.NewHandler()
	srv := server.NewServer(cfg, handlers.Init(cfg))

	go func() {
		if err := srv.Run(); !errors.Is(err, http.ErrServerClosed) {
			fmt.Println("error occurred while running http server: ", err.Error())
		}
	}()

	fmt.Println("Server started!")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)

	<-quit

	const timeout = 5 * time.Second

	ctx, shutdown := context.WithTimeout(context.Background(), timeout)
	defer shutdown()

	if err := srv.Stop(ctx); err != nil {
		fmt.Println("failed to stop server: ", err)
	}
}

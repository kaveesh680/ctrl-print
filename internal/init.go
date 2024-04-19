package internal

import (
	"context"
	"fmt"
	"github.com/kaveesh680/ctrl-print/internal/db_connection"
	"github.com/kaveesh680/ctrl-print/internal/http"
	"os"
	"os/signal"
	"syscall"
)

func Init() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	db_connection.InitConnection()
	http.InitRouter(ctx)

	select {
	case <-sigs:
		fmt.Println("sddsds")
		http.StopServer(ctx)
	}
}

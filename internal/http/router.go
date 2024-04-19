package http

import (
	"context"
	"fmt"
	"github.com/kaveesh680/ctrl-print/internal/endpoints"
	"github.com/kaveesh680/ctrl-print/internal/http_handler"
	"net/http"
	"time"
)

const prefixLog = `ctrl-print.internal.http.router`

var httpServer *http.Server

func InitRouter(ctx context.Context) {

	http.Handle("/chapter_versions/", http_handler.MakeHTTPHandler(endpoints.ChapterHistoryEndpoint, "GET"))

	startServer(ctx)

}

func startServer(ctx context.Context) {
	running := make(chan string, 1)

	httpServer = &http.Server{
		Addr:         fmt.Sprintf(`:%v`, "3000"),
		WriteTimeout: time.Second * 5,
		ReadTimeout:  time.Second * 5,
		IdleTimeout:  time.Second * 5,
		Handler:      nil,
	}

	go func(ctx context.Context) {
		err := httpServer.ListenAndServe()
		if err != nil {
			fmt.Println(
				fmt.Sprintf(
					`%v Cannot start web server : %v`,
					prefixLog,
					err,
				))
			panic(err)
		}

		running <- `done`

	}(ctx)

	<-running

	fmt.Println(
		fmt.Sprintf(
			`%v HTTP server started successfully`,
			prefixLog,
		))

}

func StopServer(ctx context.Context) {
	if err := httpServer.Shutdown(ctx); err != nil {
		fmt.Println(
			fmt.Sprintf(
				`%v Failed to gracefully shutdown server : %v`,
				prefixLog,
				err,
			))
		return
	}
}

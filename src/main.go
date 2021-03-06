package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"practice/gomicro/src/app"
	"practice/gomicro/src/app/loader"
	"practice/gomicro/src/config"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/service"
	"syscall"
	"time"
)

var wait time.Duration
var srv *http.Server

func readAppArgs() {
	configFile := flag.String("config", "../config.json",
		"Configuration file in JSON-format")

	flag.DurationVar(&wait, "graceful-timeout", time.Second*15,
		"the duration for which the server gracefully wait for existing connections to finish - e.g. 15s or 1m")

	flag.Parse()

	if len(*configFile) > 0 {
		config.FilePath = *configFile
	}
}

func loadAppData() {
	err := loader.LoadApplicationServices()
	if err != nil {
		panic(err)
	}
}

func setupReqRouter() {
	srv = lib.GetHTTPServer(service.NewRouter(), app.Service.ConfigService.ListenURL)
	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			app.Service.LoggerService.Error(err.Error())
		}
	}()
}

func handleShutDown() {
	signal.Notify(app.OSSingnal, os.Interrupt, syscall.SIGTERM)

	// Block until we receive our signal.
	<-app.OSSingnal

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	srv.Shutdown(ctx)
	// Optionally, you could run srv.Shutdown in a goroutine and block on
	// <-ctx.Done() if your application should wait for other services
	// to finalize based on context cancellation.
	app.Service.LoggerService.Error("Shuting Down")
	os.Exit(0)
}

func main() {
	fmt.Println("Starting Demo Service...!")
	readAppArgs()
	loadAppData()
	setupReqRouter()
	handleShutDown()
}

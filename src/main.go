package main

import (
	"flag"
	"fmt"
	"net/http"
	"practice/gomicro/src/app"
	"practice/gomicro/src/app/loader"
	"practice/gomicro/src/config"
	"practice/gomicro/src/lib"
	"practice/gomicro/src/service"
)

func readAppArgs() {
	configFile := flag.String("config", "../config.json",
		"Configuration file in JSON-format")
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
	middlewareManager := lib.GetMiddlewareManager()
	handler := service.NewRouter()
	middlewareManager.UseHandler(handler)
	http.ListenAndServe(app.Service.ConfigService.ListenURL, middlewareManager)
}

func main() {
	fmt.Println("Hello GoSkeleton Service!")
	readAppArgs()
	loadAppData()
	setupReqRouter()
}

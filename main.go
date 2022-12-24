package main

import (
	"os"

	"github.com/snowmerak/lux"
	"github.com/snowmerak/lux/middleware"
)

func main() {
	if len(os.Args) < 2 {
		panic("Please provide a port to listen")
	}

	port := os.Args[1]

	app := lux.New(appSwagger, middleware.SetAllowCORS)

	defer CloseDB()

	rootRoute := app.NewRouterGroup("/config")
	rootRoute.Preflight([]string{"*"}, []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, []string{"*"}, nil)
	rootRoute.GET("/:name", rootGetRouter, rootGetSwagger)
	rootRoute.POST("/:name", rootPostRouter, rootPostSwagger)
	rootRoute.PUT("/:name", rootPutRouter, rootPutSwagger)
	rootRoute.DELETE("/:name", rootDeleteRouter, rootDeleteSwagger)

	app.ShowSwagger("/swagger")

	if err := app.ListenAndServe2(":" + port); err != nil {
		panic(err)
	}
}

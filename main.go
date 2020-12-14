package main

import (
	"go-admin/routers"
	"net/http"
	"time"
)

func main() {
	r := routers.InitRouters()
	s := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    60 * time.Second,
		WriteTimeout:   60 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}

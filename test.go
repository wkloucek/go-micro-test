package main

import (
	"fmt"
	"net/http"

	httpServer "github.com/asim/go-micro/plugins/server/http/v3"
	"go-micro.dev/v4"
	"go-micro.dev/v4/server"
)

func main() {
	srv := httpServer.NewServer(
		server.Name("helloworld"),
	)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`hello world`))
	})

	hd := srv.NewHandler(mux)

	srv.Handle(hd)

	service := micro.NewService(
		micro.Server(srv),
		micro.Address("[::]:65432"),
	)
	service.Init()
	err := service.Run()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

}

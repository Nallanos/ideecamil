package main

import (
	"foundOriginalPoint/api"
	"log"
	"log/slog"

	"github.com/docker/docker/client"
)

func main() {
	client, err := client.NewClientWithOpts(client.WithHost("unix:///home/allan/.docker/desktop/docker.sock"), client.WithAPIVersionNegotiation())
	if err != nil {
		log.Fatal("Error creating docker client", err)
		panic(err)
	}

	defer client.Close()

	slog.Info("Starting API")

	app, err := api.NewAPI()
	if err != nil {
		log.Fatal("Error creating API", err)
		panic(err)
	}

	err = app.ListenAndServe()
	if err != nil {
		log.Fatal("Error starting API", err)
		panic(err)
	}

}

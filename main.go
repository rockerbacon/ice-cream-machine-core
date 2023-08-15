package main

import (
	os "os"
	rest_api "rockerbacon/ice-cream-machine-core/internal/rest_api"
	signal "os/signal"
)

func shutdownOnInterrupt(interruptionChannel <-chan os.Signal, server *rest_api.Server) {
	<-interruptionChannel

	server.Shutdown()
}

func main() {
	server := rest_api.NewServer()

	interruptionChannel := make(chan os.Signal)
	go shutdownOnInterrupt(interruptionChannel, &server)
	signal.Notify(interruptionChannel, os.Interrupt)

	server.ListenAndServe()
}

package main

import (
	log "log"
	os "os"
	rest_api "rockerbacon/ice-cream-machine-core/internal/rest_api"
	settings_reader "rockerbacon/ice-cream-machine-core/internal/settings"
	signal "os/signal"
)

func shutdownOnInterrupt(interruptionChannel <-chan os.Signal, server *rest_api.Server) {
	<-interruptionChannel

	server.Shutdown()
}

func main() {
	settings, err := settings_reader.ReadFromDefaultPath()
	if err != nil {
		log.Fatal(err)
	}

	server := rest_api.NewServer(settings.Host, settings.Port)

	interruptionChannel := make(chan os.Signal)
	go shutdownOnInterrupt(interruptionChannel, &server)
	signal.Notify(interruptionChannel, os.Interrupt)

	server.ListenAndServe()
}

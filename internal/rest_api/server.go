package rest_api

import (
	context "context"
	entrypoints "rockerbacon/ice-cream-machine-core/internal/rest_api/entrypoints"
	http "net/http"
	log "log"
	sync "sync"
)

type Server struct {
	shutdownWg sync.WaitGroup
	multiplexer *http.ServeMux
}

func NewServer() Server {
	return Server {
		multiplexer: http.NewServeMux(),
		shutdownWg: sync.WaitGroup{},
	}
}

func (self *Server)ListenAndServe() {
	// TODO panic if attempting to start twice
	self.registerEntrypoints()

	httpServer := http.Server{
			// TODO parameterize port
			Addr: "localhost:6533",
			Handler: self.multiplexer,
	}

	go httpServer.ListenAndServe()

	log.Printf("Server listening at '%s'", httpServer.Addr)

	self.shutdownWg.Add(1)
	self.shutdownWg.Wait()

	log.Println("Server shutting down")
	httpServer.Shutdown(context.Background())
}

func (self *Server)Shutdown() {
	self.shutdownWg.Done()
}

func (self *Server)registerSingleEntrypoint(e *entrypoints.Entrypoint) {
	self.multiplexer.Handle(
		e.GetPath(),
		entrypoints.NewHandler(e),
	)
}

func (self *Server)registerEntrypoints() {
	self.registerSingleEntrypoint(entrypoints.Version())
}


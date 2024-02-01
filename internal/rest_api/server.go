package rest_api

import (
	context "context"
	entrypoints "rockerbacon/ice-cream-machine-core/internal/rest_api/entrypoints"
	fmt "fmt"
	http "net/http"
	log "log"
	sync "sync"
)

type Server struct {
	host string
	multiplexer *http.ServeMux
	port uint16
	shutdownWg sync.WaitGroup
}

func NewServer(host string, port uint16) Server {
	return Server {
		host: host,
		multiplexer: http.NewServeMux(),
		port: port,
		shutdownWg: sync.WaitGroup{},
	}
}

func (self *Server)ListenAndServe() {
	// TODO panic if attempting to start twice
	self.registerEntrypoints()

	httpServer := http.Server{
		Addr: fmt.Sprintf("%s:%d", self.host, self.port),
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

func (self *Server)registerSingleEntrypoint(path string, controller any) {
	self.multiplexer.Handle(
		path,
		entrypoints.NewHandler(controller),
	)
}

func (self *Server)registerEntrypoints() {
	self.registerSingleEntrypoint("/version/", entrypoints.VersionController{})
}


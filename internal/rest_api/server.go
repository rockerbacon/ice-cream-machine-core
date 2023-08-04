package rest_api

import (
	context "context"
	http "net/http"
	log "log"
	sync "sync"
)

type Server struct {
	httpServer http.Server
	shutdownWg sync.WaitGroup
}

func NewServer() Server {
	return Server {
		httpServer: http.Server{
			// TODO parameterize port
			Addr: "localhost:6533",
			Handler: http.NewServeMux(),
		},
		shutdownWg: sync.WaitGroup{},
	}
}

func (self *Server)ListenAndServe() {
	// TODO panic if attempting to start twice
	self.httpServer.ListenAndServe()

	log.Printf("Server listening at '%s'", self.httpServer.Addr)

	self.shutdownWg.Add(1)
	self.shutdownWg.Wait()

	log.Println("Server shutting down")
	self.httpServer.Shutdown(context.Background())
}

func (self *Server)Shutdown() {
	self.shutdownWg.Done()
}


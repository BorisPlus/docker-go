package sources

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
)

type HTTPServer struct {
	server *http.Server
}

func NewHTTPServer(
	host string,
	port string,
) *HTTPServer {
	fmt.Printf("find me on address - %s:%s\n", host, port)
	mux := http.NewServeMux()
	mux.Handle("/", http.HandlerFunc(handleTeapot))
	server := http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: mux,
	}
	this := &HTTPServer{}
	this.server = &server
	return this
}

func (s *HTTPServer) Start() error {
	if err := s.server.ListenAndServe(); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	return nil
}

func (s *HTTPServer) Stop(ctx context.Context) error {
	err := s.server.Shutdown(ctx)
	if err != nil {
		return err
	}
	return nil
}

func handleTeapot(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusTeapot)
	w.Write([]byte("I receive teapot-status code!"))
}

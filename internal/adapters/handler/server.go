package handler

import (
	"context"
	"net/http"
	"port_domain_service_backend/internal/core/services"

	"github.com/sirupsen/logrus"
)

const serverPort = ":8080"

type Server struct {
	ctx         context.Context
	router      *http.ServeMux
	httpHandler *HTTPHandler
}

func AddPortDomainRoutes(m *http.ServeMux, s *Server) {
	logrus.Info("initializing routes ------ ")
	m.HandleFunc("/home", s.httpHandler.HomePage)
	m.HandleFunc("/create", s.httpHandler.CreatePortDomain)
	m.HandleFunc("/update", s.httpHandler.UpdatePortDomain)
}

func (s *Server) newRouter() {
	mux := http.NewServeMux()
	AddPortDomainRoutes(mux, s)
	s.router = mux
}

// Run starts the server and waits on completion
func (s *Server) Run() {
	l := logrus.WithField("state", "port_domain_service server")

	go func() {
		l.Infof("starting server on %s", serverPort)

		err := http.ListenAndServe(serverPort, s.router)
		if err != nil {
			l.Info(err.Error())
		}
	}()

	l.Info("main goroutine waiting for api call --> ")
loop:
	/* trunk-ignore(golangci-lint/gosimple) */
	for {
		select {
		case <-s.ctx.Done():
			l.WithError(s.ctx.Err()).Info("got ctx done")
			break loop
		}
	}
	l.Info("server run is done")
}

// NewServer runs the webserver and handles routes. Returns when the server exits.
func NewServer(ctx context.Context, ctxCancel context.CancelFunc, pds *services.PortDomainsService) (*Server, error) {
	svr := &Server{ctx: ctx}
	hdlr := NewHTTPHandler(pds)
	svr.httpHandler = hdlr
	svr.newRouter()

	return svr, nil
}

package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/mmaxim2710/firstWebApp/internal/app/config"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

type APIServer struct {
	config *config.Config
	logger *logrus.Logger
	router *mux.Router
	store  *store.Store
}

func New(config *config.Config, logger *logrus.Logger) *APIServer {
	return &APIServer{
		config: config,
		logger: logger,
		router: mux.NewRouter(),
	}
}

func (s *APIServer) Start() error {
	if err := s.configureLogger(); err != nil {
		return err
	}

	s.configureRouter()

	if err := s.configureStore(); err != nil {
		return err
	}

	s.logger.Info("Starting api server on port", s.config.Server.BindAddr)

	return http.ListenAndServe(s.config.Server.BindAddr, s.router)
}

func (s *APIServer) configureLogger() error {
	level, err := logrus.ParseLevel(s.config.LogLevel)
	if err != nil {
		return err
	}

	s.logger.SetLevel(level)

	return nil
}

func (s *APIServer) configureRouter() {
	s.router.HandleFunc("/hello", s.handleHello())
}

func (s *APIServer) configureStore() error {
	st := store.New(s.config, s.logger)
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *APIServer) handleHello() http.HandlerFunc {
	return func(rw http.ResponseWriter, req *http.Request) {
		_, err := io.WriteString(rw, "hello")
		if err != nil {
			s.logger.Error(err)
		}
	}
}

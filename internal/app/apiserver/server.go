package apiserver

import (
	"github.com/gorilla/mux"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store, logger *logrus.Logger) *server {
	s := &server{
		router: mux.NewRouter(),
		logger: logger,
		store:  store,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods("POST")
}

func (s *server) handleUsersCreate() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
	}
}

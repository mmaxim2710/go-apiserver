package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
	"github.com/sirupsen/logrus"
	"net/http"
)

type server struct {
	router *mux.Router
	logger *logrus.Logger
	store  store.Store
}

func newServer(store store.Store) *server {
	s := &server{
		router: mux.NewRouter(),
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
	type request struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	return func(w http.ResponseWriter, req *http.Request) {
		receivedReq := &request{}
		err := json.NewDecoder(req.Body).Decode(receivedReq)
		if err != nil {
			s.error(w, req, http.StatusBadRequest, err)
			return
		}

		u := &model.User{
			Email:    receivedReq.Email,
			Password: receivedReq.Password,
		}

		if err := s.store.User().Create(u); err != nil {
			s.error(w, req, http.StatusUnprocessableEntity, err)
			return
		}

		u.Sanitize()
		s.respond(w, req, http.StatusCreated, u)
	}
}

func (s *server) error(w http.ResponseWriter, req *http.Request, statusCode int, err error) {
	s.respond(w, req, statusCode, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, req *http.Request, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

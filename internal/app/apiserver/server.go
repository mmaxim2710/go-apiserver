package apiserver

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/gorilla/sessions"
	"github.com/mmaxim2710/firstWebApp/internal/app/model"
	"github.com/mmaxim2710/firstWebApp/internal/app/store"
	"net/http"
)

type server struct {
	router       *mux.Router
	store        store.Store
	sessionStore sessions.Store
}

const (
	sessionName = "moxem"
)

func newServer(store store.Store, sessionStore sessions.Store) *server {
	s := &server{
		router:       mux.NewRouter(),
		store:        store,
		sessionStore: sessionStore,
	}

	s.configureRouter()

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	s.router.ServeHTTP(w, req)
}

func (s *server) configureRouter() {
	s.router.HandleFunc("/users", s.handleUsersCreate()).Methods(http.MethodPost)
	s.router.HandleFunc("/sessions", s.handleSessionsCreate()).Methods(http.MethodPost)
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

func (s *server) handleSessionsCreate() http.HandlerFunc {
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

		u, err := s.store.User().FindByEmail(receivedReq.Email)
		if err != nil || !u.ComparePassword(receivedReq.Password) {
			s.error(w, req, http.StatusUnauthorized, ErrIncorrectEmailOrPassword)
			return
		}

		session, err := s.sessionStore.Get(req, sessionName)
		if err != nil {
			s.error(w, req, http.StatusInternalServerError, err)
			return
		}

		session.Values["user_uuid"] = u.UUID.String()
		err = s.sessionStore.Save(req, w, session)
		if err != nil {
			s.error(w, req, http.StatusInternalServerError, err)
			return
		}

		s.respond(w, req, http.StatusOK, nil)
	}
}

func (s *server) error(w http.ResponseWriter, req *http.Request, statusCode int, err error) {
	s.respond(w, req, statusCode, map[string]string{"error": err.Error()})
}

func (s *server) respond(w http.ResponseWriter, req *http.Request, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)
	if data != nil {
		_ = json.NewEncoder(w).Encode(data)
	}
}

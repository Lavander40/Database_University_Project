package server

import (
	"db_project/internal/db_project/microservices/mongo_api/store"
	"net/http"
	"github.com/gorilla/mux"
)

type Server struct {
	config *Config
	router *mux.Router
	store  *store.Store
}

func New() *Server {
	return &Server{
		router: mux.NewRouter(),
		config: NewConfig(),
	}
}

func (s *Server) Start() error {
	if err := s.configureStore(); err != nil {
		return err
	}
	s.configureRouter()

	return http.ListenAndServe(s.config.BindAddr, s.router)
}

func (s *Server) configureStore() error {
	st := store.New()
	if err := st.Open(); err != nil {
		return err
	}

	s.store = st
	return nil
}

func (s *Server) configureRouter() {
	s.router.HandleFunc("/", s.handleIndex())
}
package server

import (
	"elastic_api/store"
	"encoding/json"
	"io"
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
	s.router.HandleFunc("/lecture/{phrase}", s.handleGet()).Methods("GET")
	//s.router.HandleFunc("/set", s.handleSet()).Methods("PUT")
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "elastic api server response")
	}
}

func (s *Server) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		data, err := s.store.Get(vars["phrase"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

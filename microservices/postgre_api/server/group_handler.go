package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (s *Server) handleGroupGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		groupId, err := strconv.Atoi(vars["id"])
		if err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
			return
    	}

		data, err := s.store.Group().Get(groupId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) handleGroupGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.store.Group().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

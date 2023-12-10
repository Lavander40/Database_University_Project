package server

import (
	"encoding/json"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

func (s *Server) handleSpecialityGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		specialityId, err := strconv.Atoi(vars["id"])
		if err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
			return
    	}

		data, err := s.store.Spec().Get(specialityId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) handleSpecialityGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.store.Spec().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

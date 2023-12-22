package server

import (
	"redis_api/model"
	"encoding/json"
	"io"
	"net/http"
	"github.com/gorilla/mux"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "redis api server response")
	}
}

func (s *Server) handleGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		data, err := s.store.Student().Get(vars["id"])
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) handleSet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student model.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.store.Student().Set(student); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
	}
}

// func (s *Server) handleGetAll() http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		data, err := s.store.StudentStore.GetAll()
// 		if err != nil {
// 			io.WriteString(w, err.Error())
// 		}

// 		w.Header().Set("Content-Type", "application/json")
// 		w.WriteHeader(http.StatusCreated)
// 		json.NewEncoder(w).Encode(data)
// 	}
// }

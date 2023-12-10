package server

import (
	"db_project/internal/db_project/microservices/postgre_api/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func (s *Server) handleStudentSet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student model.Student

		err := json.NewDecoder(r.Body).Decode(&student)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if err := s.store.StudentStore.Set(student); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// body, err := json.Marshal(student)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// res, err := http.NewRequest("POST", "localhost:4041/set", bytes.NewBuffer(body))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusBadRequest)
		// 	return
		// }
		// if res.Response.StatusCode != http.StatusCreated {
		// 	http.Error(w, res.Response.Status, http.StatusBadRequest)
		// 	return
		// }

		w.WriteHeader(http.StatusCreated)
	}
}

func (s *Server) handleStudentGet() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		studId, err := strconv.Atoi(vars["id"])
		if err != nil {
        	http.Error(w, err.Error(), http.StatusBadRequest)
			return
    	}

		data, err := s.store.Student().Get(studId)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

func (s *Server) handleStudentGetAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data, err := s.store.Student().GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(data)
	}
}

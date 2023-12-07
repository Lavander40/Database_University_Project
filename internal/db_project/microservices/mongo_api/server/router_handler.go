package server

import (
	"io"
	"net/http"
)

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "postgre api server response")
	}
}

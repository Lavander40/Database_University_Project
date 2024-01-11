package server

import (
	"postgre_api/store"
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
	
	s.router.HandleFunc("/student/get", s.handleStudentGetAll()).Methods("GET")
	s.router.HandleFunc("/student/get/{id}", s.handleStudentGet()).Methods("GET")
	s.router.HandleFunc("/student/set", s.handleStudentSet()).Methods("POST")

	s.router.HandleFunc("/attend/get", s.handleAttendGetAll()).Methods("GET")
	s.router.HandleFunc("/attend/get/{id}", s.handleAttendGet()).Methods("GET")
	s.router.HandleFunc("/attend/rate", s.handleAttendRate()).Methods("GET")

	s.router.HandleFunc("/course/get", s.handleCourseGetAll()).Methods("GET")
	s.router.HandleFunc("/course/get/{id}", s.handleCourseGet()).Methods("GET")

	s.router.HandleFunc("/equip/get", s.handleEquipmentGetAll()).Methods("GET")
	s.router.HandleFunc("/equip/get/{id}", s.handleEquipmentGet()).Methods("GET")

	s.router.HandleFunc("/group/get", s.handleGroupGetAll()).Methods("GET")
	s.router.HandleFunc("/group/get/{id}", s.handleGroupGet()).Methods("GET")

	s.router.HandleFunc("/lesson/get", s.handleLessonGetAll()).Methods("GET")
	s.router.HandleFunc("/lesson/get/{id}", s.handleLessonGet()).Methods("GET")

	s.router.HandleFunc("/room/get", s.handleRoomGetAll()).Methods("GET")
	s.router.HandleFunc("/room/get/{id}", s.handleRoomGet()).Methods("GET")

	s.router.HandleFunc("/sched/get", s.handleScheduleGetAll()).Methods("GET")
	s.router.HandleFunc("/sched/get/{id}", s.handleScheduleGet()).Methods("GET")

	s.router.HandleFunc("/spec/get", s.handleSpecialityGetAll()).Methods("GET")
	s.router.HandleFunc("/spec/get/{id}", s.handleSpecialityGet()).Methods("GET")
}

func (s *Server) handleIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "postgre api server response")
	}
}

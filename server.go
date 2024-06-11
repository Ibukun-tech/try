package try

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Server struct {
	dbServices Service
}

func NewServer(serv Service) *Server {
	return &Server{
		dbServices: serv,
	}
}
func writeJsonHeader(w http.ResponseWriter, val int, v any) error {
	w.Header().Add("Content-type", "application/json")
	w.WriteHeader(val)
	return json.NewEncoder(w).Encode(v)
}

type handleFunc func(w http.ResponseWriter, r *http.Request) error

type errorApiHandle struct {
	Error error `json:"error"`
}

func RunHandler(h handleFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := h(w, r); err != nil {
			v := errorApiHandle{
				Error: err,
			}
			writeJsonHeader(w, http.StatusBadRequest, v)
		}
	}
}

func (s *Server) RegisterHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "Post" {
		var newLog Log
		if err := json.NewDecoder(r.Body).Decode(&newLog); err != nil {
			return err
		}
		value, err := s.dbServices.Add(newLog)
		if err != nil {
			return err
		}
		return writeJsonHeader(w, http.StatusCreated, value)
	}
	return nil
}
func (s *Server) GetAllHandler(w http.ResponseWriter, r *http.Request) error {
	if r.Method == http.MethodGet {
		allLog, err := s.dbServices.List()
		if err != nil {
			allLogError := &errorApiHandle{
				Error: errors.New("error in finding all logs"),
			}
			return writeJsonHeader(w, http.StatusBadGateway, allLogError)
		}
		return writeJsonHeader(w, http.StatusOK, allLog)
	}
	return nil
}

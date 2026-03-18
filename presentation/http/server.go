package presentation_http

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/S-L-T/demo-microservice/domain/entity"
	"github.com/S-L-T/demo-microservice/domain/use_case"
	"github.com/S-L-T/demo-microservice/helper"
	"net/http"
	"time"
)

type Server struct {
	Router      *mux.Router
	userUseCase use_case.User
}

func NewServer(u use_case.User) Server {
	s := Server{
		Router:      mux.NewRouter(),
		userUseCase: u,
	}
	s.initializeRoutes()
	return s
}

func (s *Server) initializeRoutes() {
	s.Router.HandleFunc("/user", s.user).Methods(
		http.MethodGet,
		http.MethodPut,
		http.MethodPatch,
		http.MethodDelete,
		http.MethodOptions,
	)
	s.Router.HandleFunc("/healthcheck", s.healthcheck).Methods(
		http.MethodGet,
		http.MethodOptions,
	)
	s.Router.Use(mux.CORSMethodMiddleware(s.Router))
}

func (s *Server) user(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		s.getHandler(w, r)
		break
	case http.MethodPut:
		s.putHandler(w, r)
		break
	case http.MethodPatch:
		s.patchHandler(w, r)
		break
	case http.MethodDelete:
		s.deleteHandler(w, r)
		break
	case http.MethodOptions:
		s.optionsHandler(w)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (s *Server) healthcheck(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func (s *Server) writeResponse(w http.ResponseWriter, resData interface{}) {
	res, err := json.Marshal(resData)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusInternalServerError)
	}

	_, err = w.Write(res)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (s *Server) getHandler(w http.ResponseWriter, r *http.Request) {
	req := entity.GetPaginatedUsersRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusBadRequest)
		resData := entity.GetPaginatedUsersResponse{
			Users: nil,
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}
	filter := entity.Filter{
		FirstName: req.Filter.FirstName,
		LastName:  req.Filter.LastName,
		Nickname:  req.Filter.Nickname,
		Email:     req.Filter.Email,
		Country:   req.Filter.Country,
	}

	users, err := s.userUseCase.GetPaginatedUsers(filter, req.PageNum, req.PageSize)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusInternalServerError)
		resData := entity.GetPaginatedUsersResponse{
			Users: nil,
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	var userResponse []entity.UserResponse
	for _, u := range users {
		userResponse = append(userResponse, entity.UserResponse{
			ID:        u.ID,
			FirstName: u.FirstName,
			LastName:  u.LastName,
			Nickname:  u.Nickname,
			Password:  u.Password,
			Email:     u.Email,
			Country:   u.Country,
			CreatedAt: u.CreatedAt.Format(time.RFC3339),
			UpdatedAt: u.UpdatedAt.Format(time.RFC3339),
		})
	}
	resData := entity.GetPaginatedUsersResponse{
		Users: userResponse,
		Error: "",
	}
	w.WriteHeader(http.StatusOK)
	s.writeResponse(w, resData)
}

func (s *Server) putHandler(w http.ResponseWriter, r *http.Request) {
	req := entity.AddUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusBadRequest)
		resData := entity.AddUserResponse{
			ID:    "",
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	user := entity.User{
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Nickname:  req.User.Nickname,
		Password:  req.User.Password,
		Email:     req.User.Email,
		Country:   req.User.Country,
	}

	id, err := s.userUseCase.AddUser(user)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusInternalServerError)
		resData := entity.AddUserResponse{
			ID:    "",
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}
	resData := entity.AddUserResponse{
		ID:    id,
		Error: "",
	}
	w.WriteHeader(http.StatusOK)
	s.writeResponse(w, resData)
}

func (s *Server) patchHandler(w http.ResponseWriter, r *http.Request) {
	req := entity.UpdateUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusBadRequest)
		resData := entity.UpdateUserResponse{
			User:  entity.UserResponse{},
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	user := entity.User{
		ID:        req.ID,
		FirstName: req.User.FirstName,
		LastName:  req.User.LastName,
		Nickname:  req.User.Nickname,
		Password:  req.User.Password,
		Email:     req.User.Email,
		Country:   req.User.Country,
	}

	err = s.userUseCase.UpdateUser(user)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusInternalServerError)
		resData := entity.UpdateUserResponse{
			User:  entity.UserResponse{},
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}
	resData := entity.UpdateUserResponse{
		Error: "",
	}
	w.WriteHeader(http.StatusOK)
	s.writeResponse(w, resData)
}

func (s *Server) deleteHandler(w http.ResponseWriter, r *http.Request) {
	req := entity.DeleteUserRequest{}
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusBadRequest)
		resData := entity.DeleteUserResponse{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}

	err = s.userUseCase.DeleteUser(req.ID)
	if err != nil {
		helper.Log(err, helper.ErrorLevel)
		w.WriteHeader(http.StatusInternalServerError)
		resData := entity.DeleteUserResponse{
			Error: err.Error(),
		}
		s.writeResponse(w, resData)
		return
	}
	resData := entity.DeleteUserResponse{
		Error: "",
	}
	w.WriteHeader(http.StatusOK)
	s.writeResponse(w, resData)
}

func (s *Server) optionsHandler(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

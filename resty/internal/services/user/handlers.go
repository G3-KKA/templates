package user

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"yet-again-templates/resty/internal/config"
	domain "yet-again-templates/resty/internal/domain/objects"
)

type /* other module */ userHandler struct {
	service IuserService
}
type IuserHandler interface {
	GetUser(w http.ResponseWriter, r *http.Request)
	CreateUser(w http.ResponseWriter, r *http.Request)
}

func NewUserHandler(ctx context.Context, config config.Config, service IuserService) (IuserHandler, error) {
	return &userHandler{
		service: service,
	}, nil
}
func (h *userHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var user struct {
		Id int `json:"id,omitempty"`
	}
	json.Unmarshal(bodybytes, &user)
	usr, err := h.service.UserById(context.TODO(), uint64(user.Id))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var uunmapper = func(mappabe *domain.User) struct {
		Id    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	} {
		return struct {
			Id    int    `json:"id,omitempty"`
			Name  string `json:"name,omitempty"`
			Email string `json:"email,omitempty"`
		}{
			Id:    int(mappabe.Id),
			Name:  mappabe.Name,
			Email: mappabe.Email,
		}
	}
	rsp, err := json.Marshal(uunmapper(usr))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "application/json")
	w.Write(rsp)

}
func (h *userHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	bodybytes, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	var user struct {
		Id    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	}
	json.Unmarshal(bodybytes, &user)
	var umapper = func(mappabe struct {
		Id    int    `json:"id,omitempty"`
		Name  string `json:"name,omitempty"`
		Email string `json:"email,omitempty"`
	}) *domain.User {
		return &domain.User{
			Id:    uint64(mappabe.Id),
			Name:  mappabe.Name,
			Email: mappabe.Email,
		}
	}

	err = h.service.CreateUser(context.TODO(), umapper(user))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}

}

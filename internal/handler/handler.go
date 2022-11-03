package handler

import (
	"github.com/Vzhrkv/avito_internship/internal/service"
	"github.com/gorilla/mux"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user/add/balance", h.AddBalance()).Methods("POST")
	router.HandleFunc("/user/get/balance", h.GetBalance()).Methods("POST")
	router.HandleFunc("/user/reserve/balance", nil).Methods("POST")
	router.HandleFunc("/user/confirm-order", nil).Methods("POST")
	return router
}

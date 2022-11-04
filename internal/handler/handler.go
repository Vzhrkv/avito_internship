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
	router.HandleFunc("/user/get/balance", h.GetBalance()).Methods("GET")
	router.HandleFunc("/user/reserve/balance", h.ReserveBalance()).Methods("POST")
	router.HandleFunc("/user/send", h.SendToOtherUser()).Methods("POST")
	router.HandleFunc("/user/confirm-order", h.ConfirmOrder()).Methods("POST")
	return router
}

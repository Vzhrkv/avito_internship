package handler

import "github.com/gorilla/mux"

type Handler struct {
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/user/add/balance", nil).Methods("POST")
	router.HandleFunc("/user/get/balance", nil).Methods("GET")
	router.HandleFunc("/user/reserve/balance", nil).Methods("POST")
	router.HandleFunc("/user/confirm-order", nil).Methods("POST")
	return router
}

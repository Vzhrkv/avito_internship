package handler

import (
	"encoding/json"
	model "github.com/Vzhrkv/avito_internship/internal/database"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) AddBalance() http.HandlerFunc {
	type input struct {
		UserID   uint `json:"user_id"`
		AddFunds uint `json:"add_funds"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		in := &input{}
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			logrus.Print(err)
		}

		u := &model.User{
			UserID:  in.UserID,
			Balance: in.AddFunds,
		}

		if err := h.service.CreateBalance(u); err != nil {
			logrus.Print(err)
		}

		h.Respond(w, r, http.StatusCreated, nil)
	}
}

func (h *Handler) GetBalance() http.HandlerFunc {
	type input struct {
		UserID uint `json:"user_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		in := &input{}
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			logrus.Print(err)
		}

		balance, err := h.service.GetBalance(in.UserID)
		if err != nil {
			logrus.Print(err)
		}
		u := &model.User{
			UserID:  in.UserID,
			Balance: balance,
		}
		h.Respond(w, r, http.StatusOK, u)

	}
}

func (h *Handler) Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		json.NewEncoder(w).Encode(data)
	}
}

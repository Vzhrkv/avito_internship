package handler

import (
	"encoding/json"
	model "github.com/Vzhrkv/avito_internship/internal/model"
	"github.com/Vzhrkv/avito_internship/logging"
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

		if err := h.service.AddBalance(in.UserID, in.AddFunds); err != nil {
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

func (h *Handler) ReserveBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := &model.Order{}
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			logrus.Print(err)
		}
		err := h.service.ReserveBalance(in.UserID, in.ServiceID, in.OrderID, in.Price)
		if err != nil {
			logrus.Print(err)
			h.Respond(w, r, http.StatusInternalServerError, nil)
		}
		h.Respond(w, r, http.StatusAccepted, nil)
	}
}
func (h *Handler) ConfirmOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := &model.Order{}
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			logrus.Print(err)
		}
		err := h.service.ConfirmOrder(in.UserID, in.ServiceID, in.OrderID, in.Price)
		if err != nil {
			logrus.Print(err)
			h.Respond(w, r, http.StatusInternalServerError, nil)
		}

		logging.LogToFile(in)

		h.Respond(w, r, http.StatusOK, nil)
	}
}

func (h *Handler) Respond(w http.ResponseWriter, r *http.Request, code int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			return
		}
	}
}

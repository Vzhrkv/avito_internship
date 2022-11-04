package handler

import (
	"encoding/json"
	"fmt"
	model "github.com/Vzhrkv/avito_internship/internal/model"
	"github.com/Vzhrkv/avito_internship/logging"
	"github.com/Vzhrkv/avito_internship/pkg/responses"
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
			res := responses.Response{
				Status: "failed",
				Msg:    "Incorrect input data",
			}
			h.Respond(w, r, http.StatusBadRequest, res)
			return
		}

		if in.UserID == 0 || in.AddFunds == 0 {
			res := responses.Response{
				Status: "failed",
				Msg:    "Can't create user with this parameters",
			}
			h.Respond(w, r, http.StatusBadRequest, res)
			return
		}

		if err := h.service.AddBalance(in.UserID, in.AddFunds); err != nil {
			logrus.Print(err)
			return
		}

		msg := fmt.Sprintf("Added funds to user(user_id=%d)", in.UserID)

		res := responses.Response{
			Status: "accepted",
			Msg:    msg,
		}

		h.Respond(w, r, http.StatusCreated, res)
	}
}

func (h *Handler) GetBalance() http.HandlerFunc {
	type input struct {
		UserID uint `json:"user_id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		in := &input{}
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			h.Respond(w, r, http.StatusBadRequest, in)
		}

		balance, err := h.service.GetBalance(in.UserID)
		if err != nil {
			res := responses.Response{
				Status: "failed",
				Msg:    "No such user",
			}
			h.Respond(w, r, http.StatusBadGateway, res)
			return
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
			res := responses.Response{
				Status: "failed",
				Msg:    err.Error(),
			}
			h.Respond(w, r, http.StatusBadRequest, res)
			return
		}
		err := h.service.ReserveBalance(in.UserID, in.ServiceID, in.OrderID, in.Price)
		if err != nil {
			res := responses.Response{
				Status: "failed",
				Msg:    err.Error(),
			}
			h.Respond(w, r, http.StatusInternalServerError, res)
			return
		}

		msg := fmt.Sprintf("Funds for order(user_id=%d, service_id=%d, order_id=%d, price=%d) was reserved",
			in.UserID, in.ServiceID, in.OrderID, in.Price)

		res := responses.Response{
			Status: "accepted",
			Msg:    msg,
		}

		h.Respond(w, r, http.StatusAccepted, res)
	}
}

func (h *Handler) ConfirmOrder() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		in := &model.Order{}
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			h.Respond(w, r, http.StatusBadRequest, in)
			return
		}
		err := h.service.ConfirmOrder(in.UserID, in.ServiceID, in.OrderID, in.Price)
		if err != nil {
			h.Respond(w, r, http.StatusInternalServerError, err)
			return
		}

		msg := fmt.Sprintf("Order(user_id=%d, service_id=%d, order_id=%d, price=%d) was confirmed",
			in.UserID, in.ServiceID, in.OrderID, in.Price)

		res := responses.Response{
			Status: "accepted",
			Msg:    msg,
		}
		logging.LogToFile(in)

		h.Respond(w, r, http.StatusOK, res)
	}
}

func (h *Handler) SendToOtherUser() http.HandlerFunc {
	type input struct {
		UserID      uint `json:"user_id"`
		OtherUserId uint `json:"other_user_id"`
		FundsToSend uint `json:"funds_to_send"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		in := &input{}
		if err := json.NewDecoder(r.Body).Decode(in); err != nil {
			h.Respond(w, r, http.StatusBadRequest, in)
		}
		err := h.service.SendToOtherUser(in.UserID, in.OtherUserId, in.FundsToSend)
		if err != nil {
			res := responses.Response{
				Status: "failed",
				Msg:    err.Error(),
			}
			h.Respond(w, r, http.StatusInternalServerError, res)
			return
		}
		res := responses.Response{
			Status: "accepted",
			Msg:    "Sended funds",
		}
		h.Respond(w, r, http.StatusOK, res)
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
		return
	}
	return
}

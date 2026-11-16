package http

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"payflow/internal/payment/domain"
	"payflow/internal/payment/service"
)

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/healthz", h.healthz).Methods(http.MethodGet)
	r.HandleFunc("/readyz", h.readyz).Methods(http.MethodGet)
	r.HandleFunc("/v1/payments", h.createPayment).Methods(http.MethodPost)
	r.HandleFunc("/v1/payments/{id}", h.getPayment).Methods(http.MethodGet)
}

func (h *Handler) healthz(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ok"})
}

func (h *Handler) readyz(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, http.StatusOK, map[string]string{"status": "ready"})
}

func (h *Handler) createPayment(w http.ResponseWriter, r *http.Request) {
	var req domain.CreatePaymentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSONError(w, http.StatusBadRequest, "invalid_json", "request body must be valid JSON")
		return
	}

	payment, err := h.service.CreatePayment(r.Context(), req)
	if err != nil {
		writeJSONError(w, http.StatusBadRequest, "validation_error", err.Error())
		return
	}

	writeJSON(w, http.StatusCreated, payment)
}

func (h *Handler) getPayment(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	payment, err := h.service.GetPayment(r.Context(), id)
	if err != nil {
		writeJSONError(w, http.StatusNotFound, "not_found", err.Error())
		return
	}

	writeJSON(w, http.StatusOK, payment)
}

func writeJSON(w http.ResponseWriter, status int, payload any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}

func writeJSONError(w http.ResponseWriter, status int, code, message string) {
	writeJSON(w, status, map[string]any{
		"error": map[string]string{
			"code":    code,
			"message": message,
		},
	})
}

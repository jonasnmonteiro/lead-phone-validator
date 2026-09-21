package handlers

import (
	"encoding/json"
	"net/http"

	"leadphone-validator/internal/models"
	"leadphone-validator/internal/validator"
)

func writeJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func HandleHealth(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodHead {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	writeJSON(w, http.StatusOK, models.HealthResponse{
		Status:  "ok",
		Service: "LeadPhone Validator",
		Version: "1.0.0",
	})
}

func HandleValidate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	var req models.ValidateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid JSON payload",
			Message: err.Error(),
		})
		return
	}

	if req.Phone == "" && req.Raw != "" {
		req.Phone = req.Raw
	}

	if req.Phone == "" {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "'phone' field is required.",
		})
		return
	}

	res := validator.Validate(req)
	writeJSON(w, http.StatusOK, res)
}

func HandleBatchValidate(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	var req models.BatchValidateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid JSON payload",
			Message: err.Error(),
		})
		return
	}

	if len(req.Phones) == 0 {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "'phones' array is required and cannot be empty.",
		})
		return
	}

	if len(req.Phones) > 10000 {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "Batch size exceeds maximum limit of 10000 items.",
		})
		return
	}

	res := validator.ValidateBatch(req.Phones)
	writeJSON(w, http.StatusOK, res)
}

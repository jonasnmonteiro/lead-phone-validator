package handlers

import (
	"encoding/json"
	"net/http"
	"strings"

	"leadphone-validator/internal/models"
	"leadphone-validator/internal/telecom"
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

func HandleDDD(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/v1/ddd/")
	path = strings.TrimPrefix(path, "/ddd/")
	path = strings.TrimPrefix(path, "/v1/ddd")
	path = strings.TrimPrefix(path, "/ddd")
	path = strings.Trim(path, "/")

	q := r.URL.Query().Get("q")
	if q == "" {
		q = r.URL.Query().Get("city")
	}
	if q == "" {
		q = r.URL.Query().Get("code")
	}
	if q == "" && path != "" {
		q = path
	}

	if q != "" && len(q) == 2 {
		if info, found := telecom.LookupDDD(q); found {
			writeJSON(w, http.StatusOK, models.DDDResponse{
				Query:   q,
				Found:   true,
				Results: []models.DDDInfo{info},
			})
			return
		}
	}

	results := telecom.SearchDDD(q)
	writeJSON(w, http.StatusOK, models.DDDResponse{
		Query:   q,
		Found:   len(results) > 0,
		Results: results,
	})
}

func HandleDDI(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	path := strings.TrimPrefix(r.URL.Path, "/v1/ddi/")
	path = strings.TrimPrefix(path, "/ddi/")
	path = strings.TrimPrefix(path, "/v1/ddi")
	path = strings.TrimPrefix(path, "/ddi")
	path = strings.Trim(path, "/")

	q := r.URL.Query().Get("q")
	if q == "" {
		q = r.URL.Query().Get("country")
	}
	if q == "" {
		q = r.URL.Query().Get("code")
	}
	if q == "" && path != "" {
		q = path
	}

	if q != "" {
		if info, found := telecom.LookupDDI(q); found {
			writeJSON(w, http.StatusOK, models.DDIResponse{
				Query:   q,
				Found:   true,
				Results: []models.DDIInfo{info},
			})
			return
		}
	}

	results := telecom.SearchDDI(q)
	writeJSON(w, http.StatusOK, models.DDIResponse{
		Query:   q,
		Found:   len(results) > 0,
		Results: results,
	})
}

func HandleNinthDigit(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	var req models.NinthDigitRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid JSON payload",
			Message: err.Error(),
		})
		return
	}

	if req.Phone == "" {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "'phone' field is required.",
		})
		return
	}

	resp := telecom.FixNinthDigit(req.Phone)
	writeJSON(w, http.StatusOK, resp)
}

func HandleWhatsAppLink(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	var req models.WhatsAppLinkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid JSON payload",
			Message: err.Error(),
		})
		return
	}

	if req.Phone == "" {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "'phone' field is required.",
		})
		return
	}

	resp := telecom.BuildWhatsAppLink(req)
	writeJSON(w, http.StatusOK, resp)
}

func HandlePatternCheck(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		writeJSON(w, http.StatusMethodNotAllowed, models.ErrorResponse{
			Error: "Method not allowed",
		})
		return
	}

	var req models.PatternCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error:   "Invalid JSON payload",
			Message: err.Error(),
		})
		return
	}

	if req.Phone == "" {
		writeJSON(w, http.StatusBadRequest, models.ErrorResponse{
			Error: "'phone' field is required.",
		})
		return
	}

	resp := telecom.CheckSuspiciousPatterns(req.Phone)
	writeJSON(w, http.StatusOK, resp)
}

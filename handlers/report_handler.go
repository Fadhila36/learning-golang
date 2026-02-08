package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"kasir-api/services"
)

type ReportHandler struct {
	service *services.TransactionService
}

func NewReportHandler(service *services.TransactionService) *ReportHandler {
	return &ReportHandler{service: service}
}

// HandleTodayReport godoc
// @Summary Get today's sales report
// @Description Get sales summary for today including total revenue, total transactions, and best selling product
// @Tags reports
// @Produce json
// @Success 200 {object} models.SalesReport
// @Failure 500 {string} string "Internal server error"
// @Router /report/hari-ini [get]
func (h *ReportHandler) HandleTodayReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	report, err := h.service.GetTodaySummary()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

// HandleReport godoc
// @Summary Get sales report by date range
// @Description Get sales summary for a specific date range. If no dates provided, returns today's report.
// @Tags reports
// @Produce json
// @Param start_date query string false "Start date (YYYY-MM-DD)"
// @Param end_date query string false "End date (YYYY-MM-DD)"
// @Success 200 {object} models.SalesReport
// @Failure 400 {string} string "Invalid date format"
// @Failure 500 {string} string "Internal server error"
// @Router /report [get]
func (h *ReportHandler) HandleReport(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	startDateStr := r.URL.Query().Get("start_date")
	endDateStr := r.URL.Query().Get("end_date")

	// If no date params, redirect to today's report
	if startDateStr == "" || endDateStr == "" {
		report, err := h.service.GetTodaySummary()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(report)
		return
	}

	// Parse dates
	startDate, err := time.Parse("2006-01-02", startDateStr)
	if err != nil {
		http.Error(w, "Invalid start_date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	endDate, err := time.Parse("2006-01-02", endDateStr)
	if err != nil {
		http.Error(w, "Invalid end_date format. Use YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	report, err := h.service.GetSummaryByDateRange(startDate, endDate)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(report)
}

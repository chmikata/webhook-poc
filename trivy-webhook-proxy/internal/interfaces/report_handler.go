package interfaces

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/chmikata/webhook-poc/trivy-webhook-proxy/internal/service"
)

type ReportHandler struct {
	service *service.Report
}

func NewReportHandler(service *service.Report) *ReportHandler {
	return &ReportHandler{service: service}
}

func (r *ReportHandler) HandleReport(writer http.ResponseWriter, request *http.Request) {
	bytes, err := os.ReadFile("input.json")
	if err != nil {
		return
	}
	var model service.VulnerabilityReport
	json.Unmarshal(bytes, &model)

	err = r.service.SendReport(model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

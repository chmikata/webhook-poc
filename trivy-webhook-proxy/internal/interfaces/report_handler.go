package interfaces

import (
	"encoding/json"
	"net/http"
	"os"

	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/domain/model"
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/usecase"
)

type ReportHandler struct {
	reportCreate *usecase.ReportCreate
}

func NewReportHandler(reportCreate *usecase.ReportCreate) *ReportHandler {
	return &ReportHandler{reportCreate: reportCreate}
}

func (r *ReportHandler) HandleReport(writer http.ResponseWriter, request *http.Request) {
	bytes, err := os.ReadFile("input.json")
	if err != nil {
		return
	}
	var model model.VulnerabilityReport
	json.Unmarshal(bytes, &model)

	err = r.reportCreate.Execute(model)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusInternalServerError)
		return
	}
}

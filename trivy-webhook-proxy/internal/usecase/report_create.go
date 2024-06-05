package usecase

import (
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/domain"
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/domain/model"
)

type ReportCreate struct {
	report *domain.Report
}

func NewReportCreate(report *domain.Report) *ReportCreate {
	return &ReportCreate{report: report}
}

func (r *ReportCreate) Execute(model model.VulnerabilityReport) error {
	return r.report.Render(model)
}

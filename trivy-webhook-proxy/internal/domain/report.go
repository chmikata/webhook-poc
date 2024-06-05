package domain

import (
	"bytes"
	"path/filepath"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/domain/model"
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/domain/repository"
)

type Report struct {
	repo     repository.ReportRepository
	template *template.Template
}

func NewReport(path string, repo repository.ReportRepository) (*Report, error) {
	t, err := template.New(filepath.Base(path)).Funcs(sprig.FuncMap()).ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return &Report{
		repo:     repo,
		template: t,
	}, nil
}

func (r *Report) CreateTitle(model model.VulnerabilityReport) (string, error) {
	return r.repo.Render(model)
}

func (r *Report) CreateBody(model model.VulnerabilityReport) (string, error) {
	return r.repo.Render(model)
}

func (r *Report) Render(model model.VulnerabilityReport) error {
	var buff bytes.Buffer
	err := r.template.Execute(&buff, model)
	if err != nil {
		return err
	}
	err = r.repo.Render(buff.String())
	if err != nil {
		return err
	}
	return nil
}

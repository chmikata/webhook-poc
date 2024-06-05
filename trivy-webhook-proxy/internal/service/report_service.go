package service

import (
	"bytes"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
)

type VulnerabilityReport struct {
	Kind     string `json:"kind"`
	Metadata struct {
		Labels struct {
			ContainerName     string `json:"trivy-operator.container.name"`
			ResourceKind      string `json:"trivy-operator.resource.kind"`
			ResourceName      string `json:"trivy-operator.resource.name"`
			ResourceNameSpace string `json:"trivy-operator.resource.namespace"`
		} `json:"labels"`
	} `json:"metadata"`
	Report struct {
		Os struct {
			Family string `json:"family"`
			Name   string `json:"name"`
		} `json:"os"`
		Summary struct {
			CriticalCount int `json:"criticalCount"`
			HighCount     int `json:"highCount"`
			MediumCount   int `json:"mediumCount"`
			LowCount      int `json:"lowCount"`
			UnknownCount  int `json:"unknownCount"`
			NoneCount     int `json:"noneCount"`
		} `json:"summary"`
		Vulnerabilities []struct {
			VulnerabilityID  string  `json:"vulnerabilityID"`
			Resource         string  `json:"resource"`
			InstalledVersion string  `json:"installedVersion"`
			FixedVersion     string  `json:"fixedVersion"`
			PublishedDate    string  `json:"publishedDate"`
			LastModifiedDate string  `json:"lastModifiedDate"`
			Severity         string  `json:"severity"`
			Title            string  `json:"title"`
			PrimaryLink      string  `json:"primaryLink"`
			Score            float64 `json:"score"`
			Target           string  `json:"target"`
		} `json:"vulnerabilities"`
	} `json:"report"`
}

type ReportRender interface {
	Render(title, body string) error
}

type Report struct {
	repo     ReportRender
	template *template.Template
}

func NewReport(path string, repo ReportRender) (*Report, error) {
	t, err := template.New(filepath.Base(path)).Funcs(sprig.FuncMap()).ParseFiles(path)
	if err != nil {
		return nil, err
	}
	return &Report{
		repo:     repo,
		template: t,
	}, nil
}

func (r *Report) SendReport(model VulnerabilityReport) error {
	title := r.createTitle(model)
	body, err := r.createBody(model)
	if err != nil {
		return err
	}
	err = r.repo.Render(title, body)
	if err != nil {
		return err
	}
	return nil
}

func (r *Report) createTitle(model VulnerabilityReport) string {
	container := model.Metadata.Labels.ContainerName
	kind := model.Metadata.Labels.ResourceKind
	name := model.Metadata.Labels.ResourceName
	namespace := model.Metadata.Labels.ResourceNameSpace
	if kind == "ReplicaSet" {
		name = name[:strings.LastIndex(name, "-")]
	}
	return fmt.Sprintf("NameSpace:%s Resource:%s Container:%s", namespace, name, container)
}

func (r *Report) createBody(model VulnerabilityReport) (string, error) {
	var buff bytes.Buffer
	err := r.template.Execute(&buff, model)
	if err != nil {
		return "", err
	}
	return buff.String(), nil
}

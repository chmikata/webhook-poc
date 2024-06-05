package infrastructure

import (
	"log"

	"github.com/chmikata/webhook-poc/trivy-webhook-proxy/internal/service"
)

var _ service.ReportRender = (*StdoutRender)(nil)

type StdoutRender struct {
}

func NewStdoutRender() *StdoutRender {
	return &StdoutRender{}
}

func (s *StdoutRender) Render(title, body string) error {
	log.Println(title)
	log.Println(body)
	return nil
}

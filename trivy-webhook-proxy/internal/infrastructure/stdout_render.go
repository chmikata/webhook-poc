package infrastructure

import (
	"fmt"

	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/domain/repository"
)

var _ repository.ReportRepository = (*StdoutRender)(nil)

type StdoutRender struct {
}

func NewStdout() *StdoutRender {
	return &StdoutRender{}
}

func (s *StdoutRender) Render(report string) error {
	fmt.Println(report)
	return nil
}

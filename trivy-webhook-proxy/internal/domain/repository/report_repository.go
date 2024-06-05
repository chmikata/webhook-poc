package repository

type ReportRepository interface {
	Render(report string) error
}

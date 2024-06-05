package cmd

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/domain"
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/infrastructure"
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/interfaces"
	"github.com/rakus-dev/sre-custom-tool/trivy-webhook-proxy/internal/usecase"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "trivy-webhook-proxy",
	Short: "Proxy trivy operator webhook",
	Long: `Proxy trivy operator webhook.

Format the contents of the VulnerabilityReport
and register it as an Issue on GitHub.`,
	RunE: func(cmd *cobra.Command, args []string) error {
		// org, _ := rootCmd.PersistentFlags().GetString("org")
		// repo, _ := rootCmd.PersistentFlags().GetString("repo")
		// token, _ := rootCmd.PersistentFlags().GetString("token")
		template, _ := cmd.PersistentFlags().GetString("template")
		fmt.Println(template)
		reportDomain, err := domain.NewReport(template, infrastructure.NewStdout())
		if err != nil {
			return err
		}
		reportUseCase := usecase.NewReportCreate(reportDomain)
		reportHandler := interfaces.NewReportHandler(reportUseCase)
		mux := http.NewServeMux()
		mux.HandleFunc("POST /report", reportHandler.HandleReport)
		server := &http.Server{
			Addr:    ":8080",
			Handler: mux,
		}
		idleConnsClosed := make(chan struct{})
		go func() {
			signalChan := make(chan os.Signal, 1)
			signal.Notify(signalChan, syscall.SIGTERM, syscall.SIGINT)
			<-signalChan
			log.Println("Signal received")
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			log.Println("Server shutdown")
			server.Shutdown(ctx)
			close(idleConnsClosed)
		}()
		log.Println("Server started")
		server.ListenAndServe()
		<-idleConnsClosed
		return nil
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// rootCmd.PersistentFlags().StringP("org", "o", "", "Organization name")
	// rootCmd.PersistentFlags().StringP("repo", "r", "", "Repository name")
	// rootCmd.PersistentFlags().StringP("token", "t", "", "Token for authentication")
	rootCmd.PersistentFlags().StringP("template", "p", "", "Template file path")

	// rootCmd.MarkPersistentFlagRequired("org")
	// rootCmd.MarkPersistentFlagRequired("repo")
	// rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("template")
}

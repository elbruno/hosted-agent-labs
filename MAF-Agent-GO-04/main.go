package main

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/microsoft/agent-framework-go/agent"
	"github.com/microsoft/agent-framework-go/provider/aguiprovider"
	"github.com/microsoft/agent-framework-go/provider/foundryprovider"
)

const (
	defaultModel = "gpt-5.4-mini"
	defaultPort  = 8088
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	if err := run(ctx, logger); err != nil {
		logger.Error("agent host stopped", "error", err)
		os.Exit(1)
	}
}

func run(ctx context.Context, logger *slog.Logger) error {
	hostedAgent, err := newAgent()
	if err != nil {
		return err
	}

	port, err := resolvePort()
	if err != nil {
		return err
	}

	mux := http.NewServeMux()
	mux.Handle("/invocations", aguiprovider.NewJSONHTTPHandler(hostedAgent, aguiprovider.HandlerConfig{
		Logger: logger,
	}))
	mux.HandleFunc("/readiness", readinessHandler)

	server := &http.Server{
		Addr:              fmt.Sprintf(":%d", port),
		Handler:           mux,
		ReadHeaderTimeout: 5 * time.Second,
	}

	serverError := make(chan error, 1)
	go func() {
		logger.Info("agent host listening", "address", server.Addr)
		serverError <- server.ListenAndServe()
	}()

	select {
	case err := <-serverError:
		if !errors.Is(err, http.ErrServerClosed) {
			return fmt.Errorf("serve agent host: %w", err)
		}
	case <-ctx.Done():
		shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			return fmt.Errorf("shut down agent host: %w", err)
		}
	}

	return nil
}

func newAgent() (*agent.Agent, error) {
	endpoint := strings.TrimSpace(os.Getenv("FOUNDRY_PROJECT_ENDPOINT"))
	if endpoint == "" {
		return nil, errors.New("FOUNDRY_PROJECT_ENDPOINT environment variable is not set")
	}

	model := strings.TrimSpace(os.Getenv("AZURE_AI_MODEL_DEPLOYMENT_NAME"))
	if model == "" {
		model = defaultModel
	}

	credential, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("create Azure credential: %w", err)
	}

	return foundryprovider.NewAgent(
		endpoint,
		credential,
		foundryprovider.ModelDeployment(model),
		foundryprovider.AgentConfig{
			Instructions: "You are a friendly assistant. Keep your answers brief.",
			Config: agent.Config{
				Name: "HelloAgent",
			},
		},
	), nil
}

func resolvePort() (int, error) {
	value := strings.TrimSpace(os.Getenv("PORT"))
	if value == "" {
		return defaultPort, nil
	}

	port, err := strconv.Atoi(value)
	if err != nil || port < 1 || port > 65535 {
		return 0, fmt.Errorf("PORT must be an integer between 1 and 65535, got %q", value)
	}
	return port, nil
}

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(`{"status":"ready"}`))
}

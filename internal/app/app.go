package app

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"

	"github.com/iamasmith/hopper/internal/config"
	"github.com/iamasmith/hopper/internal/logging"
	"github.com/iamasmith/hopper/internal/server"
	"github.com/iamasmith/hopper/internal/version"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type AppState struct {
	logger       *zap.SugaredLogger
	s            *server.ServerState
	tracer       trace.Tracer
	otelshutdown func()
}

func Setup() (*server.ServerState, *AppState) {
	app := AppState{
		logger: logging.Setup(config.Config.LogLevel),
	}
	ctx := context.Background()
	otelShutdown, err := setupOTelSDK(ctx)
	if err != nil {
		panic("FAILED OTEL SETUP")
	}
	app.tracer = otel.Tracer("hopper")
	// Handle shutdown properly so nothing leaks.
	app.otelshutdown = func() {
		err = errors.Join(err, otelShutdown(context.Background()))
	}
	app.s = server.New(ctx, config.Config.ListenBind)
	app.logger.Debug("ServerState created")

	handleFunc := func(pattern string, handlerFunc func(http.ResponseWriter, *http.Request)) {
		// Configure the "http.route" for the HTTP instrumentation.
		handler := otelhttp.WithRouteTag(pattern, http.HandlerFunc(handlerFunc))
		app.s.Mux().Handle(pattern, handler)
	}

	// Add handlers with mux here
	handleFunc("/", app.pathWalk)
	app.logger.Infof("Starting %s %s %s (Build: %s, Built: %s)", version.Name(), version.Version(), version.BuildType(), version.BuildId(), version.BuildDate())
	app.logger.Infof("Server bound to %s", config.Config.ListenBind)
	return app.s, &app
}

func (s *AppState) Stop() {
	s.otelshutdown()
	s.logger.Debug("App stopped")
}

func (a *AppState) pathWalk(w http.ResponseWriter, r *http.Request) {
	ctx := otel.GetTextMapPropagator().Extract(r.Context(), propagation.HeaderCarrier(r.Header))
	ctx, span := a.tracer.Start(ctx,
		"pathWalk",
		trace.WithSpanKind(trace.SpanKindServer),
	)
	defer span.End()
	path := r.URL.Path
	host := r.Host
	// Included leading /
	url := fmt.Sprintf("http:/%s", path)
	writef := func(s string, args ...any) {
		w.Write([]byte(fmt.Sprintf(s, args...)))
	}
	writef("Current Host: %s\n", host)
	writef("Current Request Headers...\n")
	for name, values := range r.Header {
		// Loop over all values for the name.
		writef("%s: %v\n", name, values)
	}
	if path == "/" {
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("This is the last host, no more upstream requests"))
		return
	}
	writef("Making upstream GET %s\n", url)
	writef("Reply...\n")
	// resp, err := otelhttp.Get(ctx, url)
	req, _ := http.NewRequestWithContext(
		ctx, http.MethodGet, url, nil, // http.NoBody,
	)
	httpClient := &http.Client{Transport: otelhttp.NewTransport(
		http.DefaultTransport,
	)}
	resp, err := httpClient.Do(req)
	if err != nil {
		writef("Error %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	io.Copy(w, resp.Body)
}

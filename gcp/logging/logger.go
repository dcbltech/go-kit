package logging

import (
	"context"
	"log/slog"
	"os"
)

type logger struct{ handler slog.Handler }

func newLogger(opts *slog.HandlerOptions) *logger {
	return &logger{handler: slog.NewJSONHandler(os.Stdout, opts)}
}

func (h *logger) Enabled(ctx context.Context, level slog.Level) bool {
	return h.handler.Enabled(ctx, level)
}

func (h *logger) Handle(ctx context.Context, rec slog.Record) error {
	trace := traceFromContext(ctx)
	if trace != "" {
		rec = rec.Clone()
		rec.Add("logging.googleapis.com/trace", slog.StringValue(trace))
	}

	return h.handler.Handle(ctx, rec)
}

func (h *logger) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &logger{handler: h.handler.WithAttrs(attrs)}
}

func (h *logger) WithGroup(name string) slog.Handler {
	return &logger{handler: h.handler.WithGroup(name)}
}

func traceFromContext(ctx context.Context) string {
	trace := ctx.Value("trace")
	if trace == nil {
		return ""
	}
	return trace.(string)
}

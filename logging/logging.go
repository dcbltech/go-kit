package logging

import (
	"log/slog"
	"os"
)

const LevelCritical = slog.Level(12)

func Setup(debug bool, local bool) *slog.Logger {
	var logger *slog.Logger

	o := &slog.HandlerOptions{
		AddSource: true,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if groups != nil {
				return a
			}

			switch a.Key {
			case slog.MessageKey:
				a.Key = "message"
			case slog.SourceKey:
				a.Key = "logging.googleapis.com/sourceLocation"
			case slog.LevelKey:
				a.Key = "severity"

				if level, ok := a.Value.Any().(slog.Level); ok && level == LevelCritical {
					a.Value = slog.StringValue("CRITICAL")
				}
			}

			return a
		},
	}

	if debug {
		o.Level = slog.LevelDebug
	} else {
		o.Level = slog.LevelInfo
	}

	if local {
		logger = slog.New(slog.NewTextHandler(os.Stdout, o))
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, o))
	}

	return logger
}

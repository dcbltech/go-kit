package logging

import (
	"log/slog"
	"os"
)

const LevelCritical = slog.Level(12)

func Setup(debug bool, local bool) {
	var logger *slog.Logger

	o := &slog.HandlerOptions{
		Level:     slog.LevelInfo,
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
	}

	if local {
		logger = slog.New(slog.NewTextHandler(os.Stdout, o))
	} else {
		logger = slog.New(newLogger(o))
	}

	logger.Debug("logger initialized", "debug", debug, "local", local)

	slog.SetDefault(logger)
}

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

				if level, ok := a.Value.Any().(slog.Level); ok {
					switch level {
					case LevelCritical:
						a.Value = slog.StringValue("CRITICAL")
					case slog.LevelError:
						a.Value = slog.StringValue("ERROR")
					case slog.LevelWarn:
						a.Value = slog.StringValue("WARNING")
					case slog.LevelInfo:
						a.Value = slog.StringValue("INFO")
					case slog.LevelDebug:
						a.Value = slog.StringValue("DEBUG")
					default:
						a.Value = slog.StringValue("DEFAULT")
					}
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

	logger.Debug("logger initialized", "debug", debug, "local", local)

	return logger
}

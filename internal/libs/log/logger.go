package log

import (
	"io"
	"log/slog"
	"os"
)

type Mode string

const (
	Prod Mode = "prod"
	Dev       = "dev"
)

var presets = map[Mode]slog.Level{
	Prod: slog.LevelInfo,
	Dev:  slog.LevelDebug,
}

func MustNewLogger(buildMode Mode, w io.Writer) *slog.Logger {
	if w == nil {
		panic("error to initialize production logger")
	}

	return slog.New(
		slog.NewTextHandler(w, &slog.HandlerOptions{
			Level: presets[buildMode],
		},
		),
	)
}

func MustNewLogWriter(enableConsoleLogging bool, filepath string) (logWriter io.WriteCloser) {
	if enableConsoleLogging {
		logWriter = os.Stdout
	} else {
		var err error
		logWriter, err = os.Create(filepath)
		if err != nil {
			panic("Error in opening log file")
		}
	}

	return logWriter
}

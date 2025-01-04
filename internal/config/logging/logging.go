package logging

import (
	"github.com/lallison21/library_rest_service/internal/config/config"
	"io"
	"log/slog"
	"os"
	"runtime/debug"
)

func New(cfg config.Logging) *slog.Logger {
	var out io.Writer
	if cfg.LogToFile {
		file, err := os.Create("auth_service.log")
		if err != nil {
			panic(err)
		}
		out = file
	} else {
		out = os.Stdout
	}

	handleOpt := &slog.HandlerOptions{
		AddSource: true,
	}
	if cfg.IsDebug {
		handleOpt.Level = slog.LevelDebug
	} else {
		handleOpt.Level = slog.LevelInfo
	}

	buildInfo, _ := debug.ReadBuildInfo()

	logger := slog.New(slog.NewJSONHandler(out, handleOpt)).
		With(slog.Group("program_info",
			slog.String("go_version", buildInfo.GoVersion),
			slog.String("service", cfg.LogIndex),
		))

	return logger
}

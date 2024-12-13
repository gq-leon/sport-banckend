package logger

import (
	"io"
	"log/slog"
	"os"

	slogmulti "github.com/samber/slog-multi"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/gq-leon/sport-backend/internal/adapter/config"
)

var logger *slog.Logger

func Set(cfg *config.App) {
	logger = slog.New(slog.NewTextHandler(os.Stdout, nil))

	if cfg.Env == "production" || cfg.Env == "preview" {
		logRotate := &lumberjack.Logger{
			Filename:   "log/app.log",
			MaxSize:    100,
			MaxAge:     30,
			MaxBackups: 3,
			Compress:   true,
		}

		logger = slog.New(
			slogmulti.Fanout(
				slog.NewJSONHandler(io.MultiWriter(os.Stdout, logRotate), nil),
			),
		)
	}

	slog.SetDefault(logger)
}

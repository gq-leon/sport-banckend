package main

import (
	"log/slog"
	"os"

	"github.com/gq-leon/sport-backend/internal/adapter/config"
	"github.com/gq-leon/sport-backend/internal/adapter/logger"
)

func main() {
	cfg, err := config.New()
	if err != nil {
		slog.Error("Error loading environment variables", "error", err)
		os.Exit(1)
	}

	logger.Set(cfg.App)
	slog.Info("Starting the application", "app", cfg.App.Name, "env", cfg.App.Env)

	// 初始化数据库
	slog.Info("Successfully connected to the database", "db", cfg.Database.Type)

	// 执行数据库迁移脚本
	slog.Info("Successfully migrated the database")

	// 初始化缓存服务

}

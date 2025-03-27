package log

import (
	"app1/internal/config"
	"log/slog"
	"os"
)

func New(cfg *config.Config) (Logger, error) {
	// TODO запись в файл
	// TODO проверки

	handler := slog.NewJSONHandler(os.Stdout, nil)
	return &slogWrapper{slog.New(handler)}, nil
}

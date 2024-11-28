package util

import (
	"io"
	"log/slog"
	"os"
)

func ReadFile(path string) []byte {
	logWithPath := slog.With(
		slog.String("path", path),
	)
	file, err := os.Open(path)
	if err != nil {
		logWithPath.Error("Файл не открылся")
		os.Exit(1)
	}
	defer file.Close()
	data, err := io.ReadAll(file)
	if err != nil {
		logWithPath.Error("Файл не прочитан")
		os.Exit(1)
	}
	logWithPath.Debug("Данные из файла получены")
	return data
}

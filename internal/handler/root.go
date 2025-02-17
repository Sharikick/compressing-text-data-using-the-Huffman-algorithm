package handler

import (
	"huff/internal/log"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "huff",
	Short: "Консольное приложение для сжатия текстовых данных",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	log := log.InitLog()

	log.Info("Start root command")
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Ошибка при запуске основной команды")
		os.Exit(1)
	}
}

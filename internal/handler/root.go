package handler

import (
	"huff/internal/log"
	"huff/internal/model/queue"
	"huff/internal/util/fs"
	"log/slog"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "huff",
	Short: "Консольное приложение для сжатия текстовых данных",
	Run:   run,
}

type Frequencies map[rune]int
type Codes map[rune]string

func run(cmd *cobra.Command, args []string) {
	log := log.InitLog()
	code("resources/text.txt", log)
}

func getFrequencies(data string) Frequencies {
	frequencies := make(Frequencies)
	for _, char := range data {
		frequencies[char]++
	}
	return frequencies
}

func generateCodes(node *queue.Node, code string, codes *Codes) {
	if node == nil {
		return
	}
	if node.Char != 0 {
		(*codes)[node.Char] = code
	}
	generateCodes(node.Left, code+"0", codes)
	generateCodes(node.Right, code+"1", codes)
}

func getRemains(size int) int {
	remains := size % 8

	if remains == 0 {
		return 0
	} else {
		return 8 - remains
	}
}

func code(path string, log *slog.Logger) {
	// Получаю все содержимое из файла
	data := string(fs.ReadFile(path))

	// Создаю мапу частот, где
	// ключ - это символы
	// значение - это частота появления этого символа в тексте
	frequencies := getFrequencies(data)

	queue_p := queue.CreatePriorityQueue(*log)
	for char, freq := range frequencies {
		queue_p.Add(&queue.Node{
			Frequency: freq,
			Char:      char,
		})
	}

	// Создаю дерево Хаффмана
	for queue_p.Length() > 1 {
		node_one := queue_p.PopRoot()
		node_two := queue_p.PopRoot()
		queue_p.Add(&queue.Node{
			Frequency: node_one.Frequency + node_two.Frequency,
			Left:      node_one,
			Right:     node_two,
		})
	}

	codes := make(Codes)
	generateCodes(queue_p.PopRoot(), "", &codes)

	var cipher string
	for _, char := range data {
		cipher = codes[char]
	}
	remains := getRemains(len(cipher))
	for i := 0; i < remains; i++ {
		cipher += "0"
	}

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Error")
		os.Exit(1)
	}
}

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

/*
=== Утилита cut ===

Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные

Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flags struct {
	fields    int
	delimiter string
	separated bool
}

func newFlags() *flags {
	return &flags{}
}

func parseFlags() *flags {
	flags := newFlags()

	flag.IntVar(&flags.fields, "f", 0, "'fields' - выбрать поля (колонки)")
	flag.StringVar(&flags.delimiter, "d", "\t", "'delimiter' - использовать другой разделитель")
	flag.BoolVar(&flags.separated, "s", false, "'separated' - только строки с разделителем")
	flag.Parse()

	return flags
}

func readFromStdin(f *flags) error {
	for {
		reader := bufio.NewReader(os.Stdin)
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}

		if f.separated && !strings.Contains(text, f.delimiter) {
			fmt.Println("")
			continue
		}

		splitString := strings.Split(text, f.delimiter)
		if f.fields > 0 && f.fields <= len(splitString) {
			fmt.Println("запрошенная колонка", f.fields, ": ", splitString[f.fields - 1])
			continue
		}

		return nil
	}
}


func cut() {
	flags := parseFlags()

	err := readFromStdin(flags)
	if err != nil {
		panic(err)
	}
}

func main() {
	cut()
}

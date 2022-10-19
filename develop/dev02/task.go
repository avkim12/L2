package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func unpack(s string) (str string, err error) {

	var sb strings.Builder

	for i := 0; i <= len(s)-1; i++ {
		if s[i] >= 97 && s[i] <= 122 {
			sb.WriteByte(s[i])
		} else if s[i] == '\\' {
			sb.WriteByte(s[i+1])
			i++
		} else if s[i] >= 48 && s[i] <= 57 {
			if i == 0 {
				return str, fmt.Errorf("incorrent string")
			}
			count, err := strconv.Atoi(string(s[i]))
			if err != nil {
				return str, err
			}
			sb.WriteString(strings.Repeat(string(s[i-1]), count-1))
		}
	}

	return sb.String(), nil
}

func main() {
	fmt.Println(unpack("qwe\\45"))
}

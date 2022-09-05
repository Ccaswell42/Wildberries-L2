package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
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

func UnpackString(str string) (string, error) {
	arr := []rune(str) //кастуем строку к срезу рун
	var mid string
	var res string
	if len(str) == 0 { // проверка на пустую строку
		return "", nil
	}
	if unicode.IsDigit(arr[0]) { // проверка, что первый символ не цифра
		return "", errors.New("некорректная строка")
	}
	for i, val := range arr {
		// если попадается цифра, преобразуем нашу строку соответствующим образом
		if unicode.IsDigit(val) {
			if val == '0' { // если попадается 0 - игнорируем его
				continue
			}
			if i != len(arr)-1 && unicode.IsDigit(arr[i+1]) { // если две цифры идут подряд - ошибка
				return "", errors.New("некорректная строка")
			}
			// кастуем цифру к типу int
			amount, err := strconv.Atoi(string(arr[i]))
			if err != nil {
				return "", err
			}
			// преобразуем строку c помощью функции Repeat
			mid = strings.Repeat(string(arr[i-1]), amount-1)
			res += mid
			continue
		}
		res += string(val)
	}

	return res, nil
}

func main() {
	str, err := UnpackString("ф0")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(str)
}

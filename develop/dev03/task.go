package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

/*
=== Утилита sort ===

Отсортировать строки (man sort)
Основное

Поддержать ключи

-k — указание колонки для сортировки
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Дополнительное

Поддержать ключи

-M — сортировать по названию месяца
-b — игнорировать хвостовые пробелы
-c — проверять отсортированы ли данные
-h — сортировать по числовому значению с учётом суффиксов

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

/*
	a := flag.Int("a", 0, "вспомогательное сообщение для a")
	b := flag.String("b", "lol", "вспомогательное сообщение для b")
	flag.Parse()
	fmt.Println(*a, *b)
*/

type NumString struct {
	Index int
	Str   string
}

func CheckRepeat(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

func GetNumeric(arr []string) ([]int, error) {

	var num []int
	var n string

	for _, val := range arr {
		s := []rune(val)
		if unicode.IsDigit(s[0]) && s[0] != '0' {
			for a := 0; a != len(s) && unicode.IsDigit(s[a]); a++ {
				n += string(s[a])
			}
			g, err := strconv.Atoi(n)
			if err != nil {
				return nil, err
			}
			num = append(num, g)
			n = ""
		}
	}
	sort.Ints(num)
	return num, nil
}

func HasPrefix(ful string, num []int) (bool, int) {
	for i, val := range num {
		str := strconv.Itoa(val)
		if strings.HasPrefix(ful, str) {
			return true, i
		}
	}
	return false, -1
}

func GetTail(numstr []NumString) string {

}
func SortNumeric(num []int, arr []string) []string {
	var new []string
	var tail []string
	var numstr []NumString
	for _, val := range arr {
		t, n := HasPrefix(val, num)
		if t {
			numstr = append(numstr, NumString{n, val})
		} else {
			new = append(new, val)
		}
	}
	a := 0
	for i := 0; i < len(numstr); i++ {

	}

	new = append(new, tail...)
	return new
}

func main() {
	file, err := os.Open("stroki")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var arr []string

	u := flag.Bool("u", false, "неправильное использование флага")
	r := flag.Bool("r", false, "неправильное использование флага")
	n := flag.Bool("n", false, "неправильное использование флага")
	flag.Parse()

	s := bufio.NewScanner(file)
	for s.Scan() {
		if *u {
			if CheckRepeat(arr, s.Text()) {
				continue
			}
		}

		arr = append(arr, s.Text())
	}

	sort.Strings(arr)
	if *n {
		num, err := GetNumeric(arr)
		if err != nil {
			log.Fatal(err)
		}
		arr = SortNumeric(num, arr)
	}
	if *r {
		sort.Sort(sort.Reverse(sort.StringSlice(arr)))
	}
	for _, val := range arr {
		fmt.Println(val)
	}

	GetNumeric(arr)
}

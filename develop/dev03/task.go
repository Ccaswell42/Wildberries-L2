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

// NumString - структура для работы с флагом -n
type NumString struct {
	Num int
	Str string
}

// ColumnString - структура для работы с флагом -k
type ColumnString struct {
	Substr  string
	FullStr string
}

// CheckRepeat - проферяет есть ли данная строка в массиве строк
func CheckRepeat(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

// GetNumeric подгатавливает массив строк для последующий сортировки с учетом числового значения
func GetNumeric(arr []string) ([]NumString, error) {

	var n string
	var numstr []NumString

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
			numstr = append(numstr, NumString{Num: g, Str: string(s)})

			n = ""
		}
	}
	return numstr, nil
}

// HasPrefix проверяет на наличие у строки префикса в виде определенной строки
func HasPrefix(ful string, numstr []NumString) bool {

	for _, val := range numstr {

		str := strconv.Itoa(val.Num)
		if strings.HasPrefix(ful, str) {
			return true
		}
	}

	return false
}

// Bubblesort сортирует массив структур по числовому значению методом пузырька и возвращает готовый
// отсортированный массив строк
func Bubblesort(numstr []NumString) []string {
	var arr []string
	for i := 0; i < len(numstr)-1; i++ {
		for j := len(numstr) - 1; j > i; j-- {
			if numstr[j-1].Num > numstr[j].Num {
				numstr[j].Num, numstr[j-1].Num = numstr[j-1].Num, numstr[j].Num
				numstr[j].Str, numstr[j-1].Str = numstr[j-1].Str, numstr[j].Str
			}
		}
	}
	for _, val := range numstr {
		arr = append(arr, val.Str)
	}
	return arr
}

func SortNumeric(numstr []NumString, arr []string) []string {
	var new []string
	var tail []string
	for _, val := range arr {
		if !HasPrefix(val, numstr) {
			new = append(new, val)
		}
	}
	tail = Bubblesort(numstr)

	new = append(new, tail...)
	return new
}

// PrepareColumn - подготавливает массив строк с колонками для последующей сортировки
func PrepareColumn(arr []string, k int) ([]string, []ColumnString) {
	var columnStr []ColumnString
	var sortArr []string
	var newstr []string

	for _, val := range arr {
		newstr = strings.Split(val, " ")
		if k >= len(newstr) {
			log.Fatal("sort: -k: Invalid argument")
		}
		sortArr = append(sortArr, newstr[k])
		columnStr = append(columnStr, ColumnString{FullStr: val, Substr: newstr[k]})
	}
	sort.Strings(sortArr)
	return sortArr, columnStr
}

// SearchStr - ищет соответствующую подстроку в массиве структур Columnstring, и возвращает полную строку, содержащую ее
func SearchStr(str string, columnStr []ColumnString) string {
	for _, val := range columnStr {
		if str == val.Substr {
			return val.FullStr
		}
	}
	return ""
}

// SortColumn - сортирует массив строк с колонками соотвествующим образом, и возвращает отсортированный массив.
func SortColumn(sortArr []string, columnStr []ColumnString) []string {
	var resStr []string
	for _, val := range sortArr {
		fullStr := SearchStr(val, columnStr)
		resStr = append(resStr, fullStr)
	}
	return resStr
}

func main() {
	// открывае файл указанный в аргумента
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// объвляем массив, куда построчно почитаем данные из файла
	var arr []string

	// объявляем флаги
	u := flag.Bool("u", false, "sort: -u: Invalid argument")
	r := flag.Bool("r", false, "sort: -r: Invalid argument")
	n := flag.Bool("n", false, "sort: -n: Invalid argument")
	k := flag.Int("k", 0, "sort: -k: Invalid argument")
	flag.Parse()

	// читаем из файла в наш массив
	s := bufio.NewScanner(file)
	for s.Scan() {
		// обрабатываем флаг "-u"
		if *u {
			if CheckRepeat(arr, s.Text()) {
				continue
			}
		}

		arr = append(arr, s.Text())
	}

	// обрабатываем флаг "-n"
	if *n {
		sort.Strings(arr)
		num, err := GetNumeric(arr)
		if err != nil {
			log.Fatal(err)
		}
		arr = SortNumeric(num, arr)
	}
	// обрабатываем флаг "-r"
	if *r {
		sort.Sort(sort.Reverse(sort.StringSlice(arr)))
	}
	// обрабатываем флаг "-k"
	if *k > 0 {
		sortArr, columnstr := PrepareColumn(arr, *k-1)
		arr = SortColumn(sortArr, columnstr)

	}
	// выводим результат в консоль
	for _, val := range arr {
		fmt.Println(val)
	}

}

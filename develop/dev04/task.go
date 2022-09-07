package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

// CheckAnagramms - проверяет является ли два слова анаграммами друг для друга
func CheckAnagramms(word, compareStr string) bool {
	wordArr := strings.Split(word, "")
	compareArr := strings.Split(compareStr, "")
	sort.Strings(wordArr)
	sort.Strings(compareArr)
	word = strings.Join(wordArr, "")
	compareStr = strings.Join(compareArr, "")
	if word == compareStr {
		return true
	}
	return false
}

// SearchAnagramms - ищет анаграммы в словаре для данного слова
func SearchAnagramms(str string, dicrionary []string) []string {
	var anagrammsArr []string
	for _, val := range dicrionary {
		if CheckAnagramms(str, val) {
			if str == val {
				continue
			}
			anagrammsArr = append(anagrammsArr, val)
		}
	}
	return anagrammsArr
}

// SearchRepeats - ищет повторяющиеся слова в мапе
func SearchRepeats(str string, anagramsMap map[string][]string) bool {
	for _, val := range anagramsMap {
		for _, word := range val {
			if str == word {
				return true
			}
		}
	}
	return false

}

// GetAnagramms - функция создания мапы с анаграммами
func GetAnagramms(dictionary []string) {
	anagramsMap := make(map[string][]string)
	var anagramsArr []string
	for _, val := range dictionary {
		if SearchRepeats(val, anagramsMap) {
			continue
		}
		if anagramsArr = SearchAnagramms(val, dictionary); len(anagramsArr) <= 1 {
			continue
		}
		anagramsMap[val] = anagramsArr
	}
	for key, anagrams := range anagramsMap {
		fmt.Printf("%s:\n", key)
		for _, word := range anagrams {
			fmt.Printf("\t%s\n", word)
		}
	}
}

func main() {
	// открывае файл указанный в аргумента
	file, err := os.Open(os.Args[len(os.Args)-1])
	if err != nil {
		log.Fatal()
	}
	defer file.Close()
	// объвляем массив, куда построчно почитаем данные из файла
	var dictionary []string

	// читаем из файла в наш массив
	s := bufio.NewScanner(file)
	for s.Scan() {
		dictionary = append(dictionary, s.Text())
	}
	// сразу соритруем наш массив
	sort.Strings(dictionary)

	// приводим все к Lowercase
	for i, val := range dictionary {
		dictionary[i] = strings.ToLower(val)
	}
	// создаем мапу анаграмм
	GetAnagramms(dictionary)

}

package main

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

import (
	"fmt"
	"sort"
	"strings"
)

func searchAnagram(arr []string) map[string][]string {
	mapAnagram := make(map[string][]string)
	iterationArray(arr, mapAnagram)
	deleteSingle(mapAnagram)
	sortMap(mapAnagram)
	return mapAnagram
}

func iterationArray(arr []string, mapAnagram map[string][]string) {
	for _, str := range arr {
		check := true
		strLower := strings.ToLower(str)
		for key, val := range mapAnagram {
			if determinateAnagram(key, strLower) {
				if !binarySearch(val, strLower) {
					mapAnagram[key] = append(val, strLower)
				}
				check = false
				break
			}
		}
		if check {
			mapAnagram[strLower] = append(mapAnagram[strLower], strLower)
		}
	}
}

func determinateAnagram(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	setOne := make(map[rune]int)
	for _, ch := range str1 {
		setOne[ch] += 1
	}
	setTwo := make(map[rune]int)
	for _, ch := range str2 {
		setTwo[ch] += 1
	}
	if len(setOne) != len(setTwo) {
		return false
	}
	for key, val := range setOne {
		if val != setTwo[key] {
			return false
		}
	}
	return true
}

func sortMap(mapAnagram map[string][]string) {
	for _, val := range mapAnagram {
		sort.Slice(val, func(i, j int) bool {
			return val[i] < val[j]
		})
	}
}

func deleteSingle(mapAnagram map[string][]string) {
	for key, val := range mapAnagram {
		if len(val) == 1 {
			delete(mapAnagram, key)
		}
	}
}

func binarySearch(arr []string, pattern string) bool {
	max, min := len(arr), 0
	for i := 0; i < len(arr); i++ {
		index := (max - min) / 2
		if arr[index] == pattern {
			return true
		} else if arr[index] < pattern {
			min = index
		} else {
			max = index
		}
	}
	return false
}

func main() {
	letters := []string{"слиток",
		"автобус",
		"пятка",
		"Столик",
		"Столик",
		"тяпка"}
	m := searchAnagram(letters)
	fmt.Println(m)
}

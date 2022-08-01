package main

/*
=== Утилита cut ===
Принимает STDIN, разбивает по разделителю (TAB) на колонки, выводит запрошенные
Поддержать флаги:
-f - "fields" - выбрать поля (колонки)
-d - "delimiter" - использовать другой разделитель
-s - "separated" - только строки с разделителем
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Flag struct {
	f bool
	d bool
	s bool
}

type Cut struct {
	flags     Flag
	value     int
	delimiter string
	cashe     []string
	matrix    [][]string
	filename  string
}

func (cut *Cut) CheckFlag() error {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) <= 1 {
		return errors.New("illegal option")
	}
	for i := 0; i < len(argsWithoutProg)-1; i++ {
		switch argsWithoutProg[i] {
		case "-f":
			cut.flags.f = true
			i++
			err := cut.ParceRow(argsWithoutProg[i])
			if err != nil {
				return err
			}
		case "-d":
			cut.flags.d = true
			i++
			cut.delimiter = argsWithoutProg[i]
			if i == len(argsWithoutProg)-1 {
				return errors.New("bad delimiter")
			}
		case "-s":
			cut.flags.s = true
		default:
			return errors.New("illegal option")
		}
	}
	cut.filename = argsWithoutProg[len(argsWithoutProg)-1]
	if !cut.flags.f {
		return errors.New("cut -f list [-s] [-d delim] [file ...]")
	}
	return nil
}

func (cut *Cut) ParceRow(str string) error {
	number, err := strconv.Atoi(str)
	if err != nil {
		return errors.New("illegal list value")
	}
	if number <= 0 {
		return errors.New("illegal list value")
	}
	cut.value = number - 1
	return nil
}

func (cut *Cut) ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		cut.cashe = append(cut.cashe, scanner.Text())
	}
	return nil
}

func (cut *Cut) FindIndex() {
	for _, val := range cut.cashe {
		cut.matrix = append(cut.matrix, strings.Split(val, cut.delimiter))
	}
}

func (cut *Cut) Print() {
	if len(cut.matrix[0])-1 < cut.value {
		cut.PrintEmpty()
		return
	}
	if cut.flags.s {
		cut.PrintFlagS()
		return
	}
	cut.PrintBasic()
}

func (cut *Cut) PrintEmpty() {
	for i := 0; i < len(cut.matrix); i++ {
		fmt.Println()
	}
}

func (cut *Cut) PrintFlagS() {
	for _, str := range cut.matrix {
		if len(str) < 2 {
			continue
		}
		fmt.Println(str[cut.value])
	}
}

func (cut *Cut) PrintBasic() {
	for _, str := range cut.matrix {
		fmt.Println(str[cut.value])
	}
}

func main() {
	cut := Cut{
		delimiter: "\t",
	}
	err := cut.CheckFlag()
	if err != nil {
		fmt.Println(err)
		return
	}
	err = cut.ReadFile(cut.filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	cut.FindIndex()
	cut.Print()
}

package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type FlagCash struct {
	k   bool
	row int
	n   bool
	r   bool
	u   bool
	m   bool
	c   bool
}

type Sort struct {
	flags    FlagCash
	cashe    [][]string
	filename string
}

func (sort *Sort) parcseFlagK(str string) error {
	number, err := strconv.Atoi(str)
	if err != nil {
		return errors.New("Invalid argument")
	}
	sort.flags.row = number - 1
	if sort.flags.row < 0 {
		return errors.New("Invalid argument")
	}
	return nil
}

func (sort *Sort) checkFlag() error {
	argsWithoutProg := os.Args[1:]
	sort.filename = argsWithoutProg[len(argsWithoutProg)-1]
	if len(argsWithoutProg) == 1 {
		return nil
	}
	for i := 0; i < len(argsWithoutProg)-1; i++ {
		switch argsWithoutProg[i] {
		case "-k":
			sort.flags.k = true
			i++
			err := sort.parcseFlagK(argsWithoutProg[i])
			if err != nil {
				fmt.Println(sort.flags.row)
				return errors.New("Invalid argument")
			}
		case "-n":
			sort.flags.n = true
		case "-r":
			sort.flags.r = true
		case "-u":
			sort.flags.u = true
		case "-M":
			sort.flags.m = true
		case "-c":
			sort.flags.c = true
		default:
			return errors.New("Invalid option")
		}
	}
	return nil
}

func (s *Sort) CreateCache(arr []string) {
	if s.flags.u {
		s.CreateSet(arr)
	} else {
		s.CreateMatrix(arr)
	}
}

func (sort *Sort) CreateMatrix(arr []string) {
	for _, val := range arr {
		sort.cashe = append(sort.cashe, strings.Split(val, " "))
	}
}

func (sort *Sort) CreateSet(arr []string) {
	set := make(map[string]bool)
	for _, value := range arr {
		set[value] = true
	}
	for val, _ := range set {
		sort.cashe = append(sort.cashe, strings.Split(val, " "))
	}
}

func readFile(filename string) ([]string, error) {
	var arrString []string
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	fileScanner := bufio.NewScanner(file)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		arrString = append(arrString, line)
	}
	if err := fileScanner.Err(); err != nil {
		return nil, err
	}
	file.Close()
	return arrString, nil
}

func print(arr []string) {
	for _, val := range arr {
		fmt.Println(val)
	}
}

func (s *Sort) SortBasic() {
	sort.Slice(s.cashe, func(i, j int) bool {
		if s.flags.r {
			return s.GetElement(i) > s.GetElement(j)
		}
		return s.GetElement(i) < s.GetElement(j)
	})
}

func (s *Sort) SortMonth() {
	sort.Slice(s.cashe, func(i, j int) bool {
		if s.flags.r {
			return s.GetMonth(i) > s.GetMonth(j)
		}
		return s.GetMonth(i) < s.GetMonth(j)
	})
}

func (s *Sort) SortNumber() {
	sort.Slice(s.cashe, func(i, j int) bool {
		one, _ := strconv.ParseFloat(s.GetElement(i), 64)
		second, _ := strconv.ParseFloat(s.GetElement(j), 64)
		if s.flags.r {
			return one > second
		}
		return one < second
	})
}

func (s *Sort) workingWithFlags() {
	if s.flags.m {
		s.SortMonth()
		return
	} else if s.flags.n {
		s.SortNumber()
		return
	} else {
		s.SortBasic()
	}
}

func (s *Sort) GetElement(i int) string {
	if s.flags.row < len(s.cashe[i]) {
		return s.cashe[i][s.flags.row]
	}
	return ""
}

func (s *Sort) GetMonth(i int) string {
	if s.flags.row < len(s.cashe[i]) {
		return s.DeterminateMonth(s.cashe[i][s.flags.row])
	}
	return ""
}

func (s *Sort) DeterminateMonth(month string) string {
	switch month {
	case "JAN":
		return "z1"
	case "FEB":
		return "z2"
	case "MAR":
		return "z3"
	case "APR":
		return "z4"
	case "MAY":
		return "z5"
	case "JUN":
		return "z6"
	case "JUL":
		return "z7"
	case "AUG":
		return "z8"
	case "SEP":
		return "z9"
	case "OCT":
		return "z10"
	case "NOV":
		return "z11"
	case "DEC":
		return "z12"
	default:
		return month
	}
}

func (s *Sort) EqualFile(arr []string) bool {
	val := 0
	for _, str := range s.cashe {
		for i := 0; i < len(str); i++ {
			if str[i] != arr[val] {
				return false
			}
			val += 1
		}
	}
	return true
}

func (s *Sort) Write(arr []string) {
	if s.flags.c {
		check := s.EqualFile(arr)
		if !check {
			fmt.Println("Disorder")
		}
	}
	s.Print()
}

func (s *Sort) Print() {
	for _, arr := range s.cashe {
		for i := 0; i < len(arr); i++ {
			if i == len(arr)-1 {
				fmt.Printf("%s", arr[i])
				break
			}
			fmt.Printf("%s ", arr[i])
		}
		fmt.Println()
	}
}

func main() {
	sort := Sort{}
	err := sort.checkFlag()
	if err != nil {
		fmt.Println(err)
		return
	}
	arr, err := readFile(sort.filename)
	if err != nil {
		fmt.Println(err)
	}
	sort.CreateCache(arr)
	sort.workingWithFlags()
	sort.Write(arr)
}

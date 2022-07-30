package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func downloadFile(filepath string, url string) error {
	/*
		Создание файла
	*/
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()
	/*
		Получить данные по ссылке
	*/
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	/*
		Проверка получение данных клиентом
	*/
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}
	/*
		Запись страницы в файл
	*/
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	argsWithoutProg := os.Args[1:]
	if len(argsWithoutProg) != 1 {
		fmt.Println("unable to resolve address")
	}
	strs := strings.Split(argsWithoutProg[0], "/")
	err := downloadFile(strs[len(strs)-1], argsWithoutProg[0])
	if err != nil {
		fmt.Println(err)
	}
}

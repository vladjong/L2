package main

import (
	"fmt"
	"sync"
	"time"
)

func or(cs ...<-chan interface{}) <-chan interface{} {
	/*
		Создаем single-chanel
	*/
	out := make(chan interface{})
	var wg sync.WaitGroup
	wg.Add(len(cs))
	/*
		Перебираем все наши каналы
	*/
	for _, c := range cs {
		/*
			Создаем новую горутину, которая будет записывать данные в single-chanel
		*/
		go func(c <-chan interface{}) {
			out <- c
			wg.Done()
		}(c)
	}
	/*
		Создаем новую горутину, которая будет дожидаться выполнения канала и закрывать его
	*/
	go func() {
		wg.Wait()
		close(out)
	}()
	return out
}

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}
	start := time.Now()
	<-or(
		sig(2*time.Second),
		sig(5*time.Second),
		sig(1*time.Second),
		sig(3*time.Second),
		sig(4*time.Second),
	)
	fmt.Printf("fone after %v", time.Since(start))
}

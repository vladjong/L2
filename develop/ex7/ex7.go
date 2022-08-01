package main

/*
=== Or channel ===
Реализовать функцию, которая будет объединять один или более done каналов в single канал если один из его составляющих каналов закроется.
Одним из вариантов было бы очевидно написать выражение при помощи select, которое бы реализовывало эту связь,
однако иногда неизестно общее число done каналов, с которыми вы работаете в рантайме.
В этом случае удобнее использовать вызов единственной функции, которая, приняв на вход один или более or каналов, реализовывала весь функционал.
Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}
Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
}()
return c
}
start := time.Now()
<-or (
	sig(2*time.Hour),
	sig(5*time.Minute),
	sig(1*time.Second),
	sig(1*time.Hour),
	sig(1*time.Minute),
)
fmt.Printf(“fone after %v”, time.Since(start))
*/

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

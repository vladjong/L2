package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

/*
	Стуктура клиента
*/
type Client struct {
	host    string
	port    string
	timeout time.Duration
	conn    net.Conn
}

/*
	Конструктор клиента
*/
func NewClient(host string, port string, timeout time.Duration) *Client {
	return &Client{
		host:    host,
		port:    port,
		timeout: timeout,
	}
}

/*
	Подключение к адрессу сервера по адрессу сети до истечения timeout.
	Используем сеть для связи TCP
	Возращаемый параметр conn - интерфейс net.Conn, который можно использовать, как потока вывода или записи.
*/
func (client *Client) ConnectionTCP() error {
	connectStr := fmt.Sprintf("%s:%s", client.host, client.port)
	fmt.Println("Trying", connectStr, "...")
	conn, err := net.DialTimeout("tcp", connectStr, client.timeout)
	if err != nil {
		return err
	}
	fmt.Println("Connected to", connectStr, ".")
	client.conn = conn
	return nil
}

/*
	Получение данных запроса. Используем чтенение через буффер Reader.
	Считывем строку методом ReadString(), пока не встрети '\n'
	Если пользователь вводит 'ctr+D' выходи из программы.
	Далее, получаем ответ от сервера.

*/
func (client *Client) sendRequest() {
	reader := bufio.NewReader(os.Stdin)
	serverReader := bufio.NewReader(client.conn)
	clientRequest, err := reader.ReadString('\n')
	if err != nil {
		if err == io.EOF {
			return
		}
		fmt.Println(err)
		return
	}
	_, err = fmt.Fprintf(client.conn, clientRequest)
	for {
		response, err := serverReader.ReadString('\n')
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Print(response)
	}
}

func main() {
	timeoutValue := flag.Uint("timeout", 10, "conn timeout")
	timeout := time.Duration(*timeoutValue) * time.Second
	flag.Parse()
	host := flag.Arg(0)
	port := flag.Arg(1)
	client := NewClient(host, port, timeout)
	err := client.ConnectionTCP()
	if err != nil {
		fmt.Println(err)
		return
	}
	client.sendRequest()
}

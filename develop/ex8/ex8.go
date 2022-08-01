package main

/*
=== Взаимодействие с ОС ===
Необходимо реализовать собственный шелл
встроенные команды: cd/pwd/echo/kill/ps
поддержать fork/exec команды
конвеер на пайпах
Реализовать утилиту netcat (nc) клиент
принимать данные из stdin и отправлять в соединение (tcp/udp)
Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

type Shell struct {
	dir string
}

func main() {
	shell := Shell{
		dir: "%",
	}
	for {
		fmt.Printf("%s ", shell.dir)
		commands := parceArgs()
		for _, command := range commands {
			if strings.Index(command, "cd") == 0 {
				shell.cd(command)
			} else if strings.Index(command, "pwd") == 0 {
				shell.pwd()
			} else if strings.Index(command, "echo") == 0 {
				shell.echo(command)
			} else if strings.Index(command, "kill") == 0 {
				shell.kill(command)
			} else if strings.Index(command, "ps") == 0 {
				shell.ps(command)
			} else if strings.Index(command, "\\quit") == 0 ||
				strings.Index(command, "\\q") == 0 {
				return
			} else {
				fmt.Println("command not found:", command)
			}
		}
	}
}

func parceArgs() []string {
	var scanner = bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		return strings.Split(scanner.Text(), " | ")
	}
	return nil
}

func (shell *Shell) cd(command string) {
	args := command[len("cd")+1:]
	os.Chdir(args)
	newDir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	shell.dir = newDir + " %"
}

func (shell *Shell) pwd() {
	path, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(path)
}

func (shell *Shell) echo(command string) {
	args := command[len("echo")+1:]
	fmt.Println(args)
}

func (shell *Shell) kill(command string) {
	args := command[len("echo")+1:]
	pid, _ := strconv.Atoi(args)
	err := syscall.Kill(pid, 9)
	if err != nil {
		fmt.Println(err)
	}
}

func (shell *Shell) ps(str string) {
	out, err := exec.Command(str).Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))
}

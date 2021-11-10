package main

import (
	"bufio"
	"fmt"
	"github.com/atotto/clipboard"
	"log"
	"os"
	"strings"
)

func main() {
	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) != os.ModeNamedPipe {
		log.Fatal("The command is intended to work with pipes.")
	}
	inputStringSlice := []string{}
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		inputStringSlice = append(inputStringSlice, s.Text())
		inputString := strings.Join(inputStringSlice, "\n")
		clipboard.WriteAll(inputString)
		text, _ := clipboard.ReadAll()
		fmt.Println(text)
	}
}

package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"os"
)

func main() {
	//cmd := exec.Command("pwd")
	//cmdOutBytes, _ := cmd.CombinedOutput()
	//cmdOut := string(cmdOutBytes)
	fileCmd,_ := os.Getwd()
	clipboard.WriteAll(fileCmd)
	text, _ := clipboard.ReadAll()
	fmt.Println(text)
	fmt.Println("Path Copyed")
}

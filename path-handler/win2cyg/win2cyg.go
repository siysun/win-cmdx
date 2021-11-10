package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func winPathTowinCYGPath(winPathString string) {
	tempString := strings.Replace(winPathString, "\\", "/", -1)
	if !strings.Contains(tempString, ":") {
		fmt.Println("Input path is not windows path")
	} else {
		splitedPathStringSlice := strings.Split(tempString, ":")
		diskNameSlice := []string{"/cygdrive/"}
		cygPathSlice := append(diskNameSlice, splitedPathStringSlice...)
		cygPathString := strings.Join(cygPathSlice, "")
		fmt.Printf("%s\n", cygPathString)
	}
}


func main() {
	/*
		Convert windows path to windows cygdrive path
		input one windows path string via arg[1] or pipe stdin
		out put windows cygpath path string
		input example:
		1. pwd|win2cyg
		2. win2cyg C:\Windows\System32\drivers\etc
	*/

	fileInfo, _ := os.Stdin.Stat()
	if (fileInfo.Mode() & os.ModeNamedPipe) == os.ModeNamedPipe {
		//Use pipe input
		inputStringSlice := []string{}
		s := bufio.NewScanner(os.Stdin)
		for s.Scan() {
			inputStringSlice = append(inputStringSlice, s.Text())
		}
		switch argLength := len(inputStringSlice); argLength {
		case 1:
			winPathString := inputStringSlice[0]
			winPathTowinCYGPath(winPathString)
		default:
			fmt.Println("Please input one win path")
		}

	} else {
		//Use arg input
		switch argLength := len(os.Args); argLength {
		case 2:
			winPathString := os.Args[1]
			winPathTowinCYGPath(winPathString)
		default:
			fmt.Println("Please input one win path")
		}
	}

}

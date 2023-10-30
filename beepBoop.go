package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func main() {
	fmt.Println("Waiting for input...")
	reader := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter text: ")
	for reader.Scan() {
		isNumber, err := regexp.MatchString(`[0-9]+`, reader.Text())
		if err != nil {
			fmt.Println("Error: ", err)
		}
		if isNumber {
			number, err := strconv.Atoi(reader.Text())
			if err != nil {
				fmt.Println("Error: ", err)
			}
			osa, err := exec.LookPath("osascript")
			for i := 0; i < number; i++ {
				cmd := exec.Command(osa, "-e", `beep`)
				cmd.Run()
			}
		}
	}
}

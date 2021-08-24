package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	fmt.Println("Hello, wkhtmltopdf")
	cmd := exec.Command("wkhtmltopdf", "https://www.douban.com/", "douban.pdf")

	var out bytes.Buffer
	cmd.Stdout = &out

	// set stderr to the provided writer, or create a new buffer
	err := cmd.Run()
	if err != nil {
		log.Fatal("run error", err)
	}

	state := cmd.ProcessState.String()
	success := cmd.ProcessState.Success()
	fmt.Println("输出:", out.String())
	fmt.Println("状态:", state)
	fmt.Println("成功:", success)
}

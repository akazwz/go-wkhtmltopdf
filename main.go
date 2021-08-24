package main

import (
	"fmt"
	"go-wkhtmltopdf/pdf"
	"log"
)

func main() {
	fmt.Println("Hello, wkhtmltopdf")
	err, filepath := pdf.GeneratePdfFromURL("https://www.douban.com/")
	if err != nil {
		log.Fatalln("generate error:", err)
	}

	fmt.Println("filepath:", filepath)
}

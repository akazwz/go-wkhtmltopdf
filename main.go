package main

import (
	"fmt"
	"go-wkhtmltopdf/image"
	"go-wkhtmltopdf/pdf"
	"log"
)

func main() {
	fmt.Println("Hello, wkhtmltopdf")
	err, filepath := pdf.GeneratePdfFromURL("https://www.douban.com/")

	if err != nil {
		log.Fatalln("generate error:", err)
	}

	fmt.Println("pdf filepath:", filepath)

	err, filepath = image.GenerateImageFromURL("https://www.douban.com/")
	if err != nil {
		log.Fatalln("generate error:", err)
	}
	fmt.Println("image filepath:", filepath)
}

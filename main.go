package main

import (
	"fmt"
	"go-wkhtmltopdf/image"
	"go-wkhtmltopdf/pdf"
	"log"
)

func main() {
	fmt.Println("Hello, wkhtmltopdf")

	err, filepath := pdf.GeneratePdfFromURL("https://s.weibo.com/top/summary/", "public/")

	if err != nil {
		log.Fatalln("generate error:", err)
	}

	fmt.Println("pdf filepath:", filepath)

	err, filepath = image.GenerateImageFromURL("https://s.weibo.com/top/summary/", "public/")
	if err != nil {
		log.Fatalln("generate error:", err)
	}
	fmt.Println("image filepath:", filepath)
}

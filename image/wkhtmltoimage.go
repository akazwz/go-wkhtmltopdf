package image

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"log"
	"os/exec"
)

func GenerateImageFromURL(url string) (err error, filepath string) {
	fileName := uuid.NewV4().String()
	filepath = fileName + ".png"
	cmd := exec.Command("wkhtmltoimage", url, filepath, "--javascript-delay", "3000")

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	log.Println("cmd:", cmd.String())

	log.Println("stderr:", stderr.String())
	// set stderr to the provided writer, or create a new buffer
	err = cmd.Run()
	if err != nil {
		log.Fatalln("run error", err)
	}

	return
}

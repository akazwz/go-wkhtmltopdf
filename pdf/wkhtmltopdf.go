package pdf

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"log"
	"os/exec"
)

func GeneratePdfFromURL(url string) (err error, filepath string) {
	fileName := uuid.NewV4().String()
	filepath = fileName + ".pdf"
	cmd := exec.Command("wkhtmltopdf", url, filepath, "--javascript-delay", "3000")

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	log.Println("stderr:", stderr.String())
	// set stderr to the provided writer, or create a new buffer
	err = cmd.Run()
	if err != nil {
		log.Fatalln("run error", err)
	}

	return
}

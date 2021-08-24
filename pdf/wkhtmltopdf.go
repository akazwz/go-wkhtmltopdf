package pdf

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"log"
	"os/exec"
	"path/filepath"
)

// GeneratePdfFromURL
// url
func GeneratePdfFromURL(url string, path string) (err error, file string) {
	fileName := uuid.NewV4().String()
	file = path + fileName + ".pdf"

	file, err = filepath.Abs("file")
	if err != nil {
		log.Fatalln("file path error")
	}

	cmd := exec.Command("wkhtmltopdf", "--javascript-delay", "3000", url, file)

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

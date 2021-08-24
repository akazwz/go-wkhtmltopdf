package image

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"go-wkhtmltopdf/utils"
	"log"
	"os/exec"
	"path/filepath"
)

// GenerateImageFromURL
// url
// image type such as jpg (default) png
func GenerateImageFromURL(url string, path string, arg ...string) (err error, file string) {
	fileName := uuid.NewV4().String()
	imageType := "jpg"
	if len(arg) >= 1 {
		imageType = arg[0]
	}

	err = utils.PathExistedAndCreate(path)
	if err != nil {
		log.Fatal("path error")
	}

	file = path + fileName + "." + imageType

	file, err = filepath.Abs(file)
	if err != nil {
		log.Fatalln("file path error")
	}
	cmd := exec.Command("wkhtmltoimage", "--javascript-delay", "3000", url, file)

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

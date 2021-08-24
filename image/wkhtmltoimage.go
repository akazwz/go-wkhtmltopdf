package image

import (
	"bytes"
	uuid "github.com/satori/go.uuid"
	"log"
	"os/exec"
)

// GenerateImageFromURL
// url
// image type such as jpg (default) png
func GenerateImageFromURL(url string, path string, arg ...string) (err error, filepath string) {
	fileName := uuid.NewV4().String()
	imageType := "jpg"
	if len(arg) >= 1 {
		imageType = arg[0]
	}
	filepath = path + fileName + "." + imageType
	cmd := exec.Command("wkhtmltoimage", "--javascript-delay", "3000", url, filepath)

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

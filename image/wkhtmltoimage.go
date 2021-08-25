package image

import (
	"bytes"
	"github.com/akazwz/go-wkhtmltopdf/utils"
	uuid "github.com/satori/go.uuid"
	"os/exec"
	"path/filepath"
	"time"
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
		return err, ""
	}

	location, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		return err, ""
	}
	now := time.Now().In(location)

	timeStr := now.Format("2006-01-02-15-04-05")

	file = path + timeStr + "-" + fileName + "." + imageType

	file, err = filepath.Abs(file)
	if err != nil {
		return err, ""
	}
	cmd := exec.Command("wkhtmltoimage", "--javascript-delay", "3000", url, file)

	var out bytes.Buffer
	cmd.Stdout = &out

	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	// set stderr to the provided writer, or create a new buffer
	err = cmd.Run()
	if err != nil {
		return err, ""
	}
	return
}

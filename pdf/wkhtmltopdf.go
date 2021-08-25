package pdf

import (
	"bytes"
	"github.com/akazwz/go-wkhtmltopdf/utils"
	uuid "github.com/satori/go.uuid"
	"os/exec"
	"path/filepath"
	"time"
)

// GeneratePdfFromURL
// url
// path
func GeneratePdfFromURL(url string, path string) (err error, file string) {
	fileName := uuid.NewV4().String()

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

	file = path + timeStr + "-" + fileName + ".pdf"
	file, err = filepath.Abs(file)
	if err != nil {
		return err, ""
	}

	cmd := exec.Command("wkhtmltopdf", "--javascript-delay", "3000", url, file)

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

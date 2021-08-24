package utils

import (
	"os"
)

func PathExistedAndCreate(path string) (err error) {
	_, err = os.Stat(path)
	if err != nil {
		err = os.Mkdir(path, os.ModePerm)
		if err != nil {
			return
		}
	}
	return
}

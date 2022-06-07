package utils

import (
	"fmt"
	"io/ioutil"
)

func ReadFile(fileLocation string) []byte {
	file, err := ioutil.ReadFile(fileLocation)
	if err != nil {
		fmt.Println(err)
	}

	return file
}

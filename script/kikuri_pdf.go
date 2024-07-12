package script

import (
	"fmt"
	"io/ioutil"
)

func KikuriPdf() string {
	filename := "data/kikuri-namu.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	}

	return string(content)
}

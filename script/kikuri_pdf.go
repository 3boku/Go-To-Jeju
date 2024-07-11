package script

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func KikuriPdf() string {
	var filteredScript string
	pattern := `（きくり）([^（|♪]*)`
	r := regexp.MustCompile(pattern)

	filename := "data/kikuri-namu.txt"
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
	
		text := string(content)
		matches := r.FindAllStringSubmatch(text, -1)
		for _, match := range matches {
			if len(match) > 1 {
				filteredScript += match[1]
			}
		}
		filteredScript += "\n"
	}

	return filteredScript
}

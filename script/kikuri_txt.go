package script

import (
	"fmt"
	"io/ioutil"
	"regexp"
)

func KikuriTxt() string {
	var filteredScript string
	pattern := `（きくり）([^（|♪]*)`
	r := regexp.MustCompile(pattern)

	for i := 1; i <= 12; i++ {
		filename := fmt.Sprintf("data/S%02d.txt", i)
		content, err := ioutil.ReadFile(filename)
		if err != nil {
			fmt.Println("Error reading file:", err)
			continue
		}

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

package script

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	var all_txt string

	for i := 1; i <= 12; i++ {
		var fileName string
		if i < 10 {
			fileName = "script/S0" + strconv.Itoa(i) + ".srt"
		} else {
			fileName = "script/S" + strconv.Itoa(i) + ".srt"
		}

		file, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			all_txt += scanner.Text() + "\n"
		}
		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}
	}

	writeToFile("output.txt", all_txt)
	fmt.Println("Text extracted and saved to output.txt")
}

func writeToFile(filename string, data string) {
	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString(data)
	if err != nil {
		log.Fatal(err)
	}
}

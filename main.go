package main

import (
	"bufio"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

var (
	filePath string
)

func init() {
	flag.StringVar(&filePath, "file", "", "reading from file")
}

func randEmoji() string {
	start := 0x1F600 // Start of emoji range
	end := 0x1F64F   // End of emoji range
	randomEmoji := rand.Intn(end-start+1) + start
	return fmt.Sprintf("%c", randomEmoji)
}

func main() {
	flag.Parse()

	var content string
	if filePath != "" {
		fmt.Println("enter a filePath")
		bytes, err := os.ReadFile(filePath)
		if err != nil {
			panic(err)
		}
		content = string(bytes)
	} else {
		scn := bufio.NewScanner(os.Stdin)
		fmt.Println("Введите текст (ctrl+] для выхода)")
		var lines []string
		for scn.Scan() {
			line := scn.Text()
			if len(line) == 1 {
				if byte(line[0]) == 4 {
					break
					
				}
			}
			lines = append(lines, line)
		}
		content = strings.Join(lines, "\n")
	}

	words := strings.Split(string(content), " ")

	out := ""

	for _, w := range words {
		out += fmt.Sprintf("%s %s ", w, randEmoji())
	}

	fmt.Println(out)
}

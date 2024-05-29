package fileops

import (
	"bufio"
	"os"
	"strings"
)

func ReadUserInput(prompt string) string {
	print(prompt + " ")
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}

	return strings.TrimSpace(text)
}

package getInput

import (
	"bufio"
	"os"
	"strings"
)

func ReadUserInput() string {
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	
	return strings.TrimSpace(text)
	
}
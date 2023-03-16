package reader

import (
	"bufio"
	"os"
)

// InputReader use bufio Package read from console and save into string
func InputReader() string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		return ""
	}

	return scanner.Text()
}

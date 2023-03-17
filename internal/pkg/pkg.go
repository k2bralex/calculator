package pkg

import (
	"bufio"
	"os"
)

func Contains(r rune, sl []rune) bool {
	for _, v := range sl {
		if v == r {
			return true
		}
	}
	return false
}

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

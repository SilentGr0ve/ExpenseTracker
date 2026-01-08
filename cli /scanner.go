package cli_

import (
	"bufio"
	"os"
)

func userInput() string {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {

	}
}

package cli

import (
	"fmt"
)

func PrintHelp() {
	fmt.Println("Commands:")
	fmt.Printf("\t1. help \n")
	fmt.Printf("\t2. add {\n")
	fmt.Printf("\t3. update \n")
	fmt.Printf("\t4. delete \n")
	fmt.Printf("\t5. list \n")
	fmt.Printf("\t6. summary \n")
}

func Current() {
	fmt.Println("Current expenses:")
}

func PrintList() {
}

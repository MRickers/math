package display

import (
	"fmt"
)

func printLines() {
	line := ""

	for i := 0; i < 15; i++ {
		line += "-"
	}
	fmt.Println(line)
}

func PrintHello(text string) {
	fmt.Println("Hello from %s", text)
	printLines()
}

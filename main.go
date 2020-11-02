package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	items := &TransactionList{}
	cmdClear()
	cmdHelp()

	for {
		fmt.Print("> ")
		commands(get(reader), items)
	}
}

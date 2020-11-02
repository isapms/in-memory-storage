package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	CmdRead   = "read"
	CmdWrite  = "write"
	CmdDelete = "delete"
	CmdStart  = "start"
	CmdCommit = "commit"
	CmdAbort  = "abort"
	CmdHelp   = "help"
	CmdClear  = "clear"
	CmdQuit   = "quit"
)

func get(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func recoverCommand(text string) {
	if r := recover(); r != nil {
		fmt.Fprintln(os.Stderr, "Unknown command: ", text)
	}
}

func cmdHelp() {
	fmt.Println("> K/V REPL with nested transactions ")
	fmt.Println("> Available commands: ")
	fmt.Println("> READ <key>			- Reads and prints to stdout, the val associated with key. If the value is not present an error is printed to stderr. ")
	fmt.Println("> WRITE <key> <val>		- Stores val in key. ")
	fmt.Println("> DELETE <key>			- Removes all key from store. Future READ commands on that key will return an error. ")
	fmt.Println("> START				- Start a transaction. ")
	fmt.Println("> COMMIT			- Commit a transaction. All actions in the current transaction are committed to the parent tx or the root store. If there is no current tx an error is output to stderr. ")
	fmt.Println("> ABORT				- Abort a transaction. All actions in the current transaction are discarded. ")
	fmt.Println("> HELP				- Show you the Help menu. ")
	fmt.Println("> CLEAR				- Clear the terminal screen. ")
	fmt.Println("> QUIT				- Exit the REPL cleanly. A message to stderr may be output. ")
	fmt.Println("> ")
}

func cmdClear() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func cmdExit() {
	fmt.Println("REPL terminated")
	os.Exit(0)
}

func commands(input string, items *TransactionList) {
	defer recoverCommand(input)

	var err error
	var result string
	operation := strings.Fields(input)

	switch strings.ToLower(operation[0]) {
	case CmdRead:
		result, err = items.Get(operation[1])

	case CmdWrite:
		items.Set(operation[1], operation[2])

	case CmdDelete:
		items.Delete(operation[1])

	case CmdStart:
		items.Begin()

	case CmdAbort:
		err = items.Rollback()

	case CmdCommit:
		err = items.Commit()

	case CmdHelp:
		cmdHelp()
		return

	case CmdClear:
		cmdClear()
		return

	case CmdQuit:
		cmdExit()

	default:
		fmt.Fprintln(os.Stderr, "Unrecognised Operation: ", input)
		return
	}

	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	if result != "" {
		fmt.Println(result)
		return
	}

	fmt.Println("OK")
}

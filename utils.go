package main

import (
	"errors"
	"fmt"
)

const (
	NoActiveTransactions ErrorMsg = "No Active Transactions"
	KeyNotFound          ErrorMsg = "Key not found: %s"
)

type ErrorMsg string

func (e ErrorMsg) Error(values ...interface{}) error {
	return errors.New(fmt.Sprintf(string(e), values...))
}

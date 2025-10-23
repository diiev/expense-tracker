package cmd

import (
	"fmt"
	"os"
)

func Run() {
	if len(os.Args) < 2 {
		showUsgae()
		return
	}
	command := os.Args[1]
	switch command {
	case "list":
		cli.ShowExpense()
	}
}

func showUsgae() {
	fmt.Println("Show Usage")
}

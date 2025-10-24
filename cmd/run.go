package cmd

import (
	"fmt"
	"os"
	"strconv"

	"go.mod/internal/cli"
)

func Run() {
	if len(os.Args) < 2 {
		showUsgae()
		return
	}
	command := os.Args[1]
	switch command {
	case "list":
		cli.ShowExpense("")
	case "add":
		amount, _ := strconv.Atoi(os.Args[2])
		amountf := float64(amount)
		cli.AddExpense(amountf, os.Args[3], os.Args[4])
	}

}

func showUsgae() {
	fmt.Println("Show Usage")
}

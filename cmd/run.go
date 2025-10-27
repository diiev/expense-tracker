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
		fmt.Printf("ID |\t  Дата\t   |\tКатегория |\tОписание\t|\tСумма\t\t|\n")
		cli.ShowExpense("")
	case "add":
		amount, _ := strconv.Atoi(os.Args[4])
		amountf := float64(amount)
		cli.AddExpense(os.Args[2], os.Args[3], amountf)
	case "delete":
		id, _ := strconv.Atoi(os.Args[2])
		cli.DeleteExpense(id)
	case "summary":
		if len(os.Args) > 2 {
			month, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Println("Ошибка преобразования аргумента в число:", err)
				return
			}
			cli.SummaryExp(month)
		} else {
			cli.SummaryExp(0)
		}
	case "update":
		id, _ := strconv.Atoi(os.Args[2])
		if len(os.Args) > 5 && len(os.Args) < 7 {
			amount, _ := strconv.Atoi(os.Args[5])
			amountf := float64(amount)
			cli.UpdateExp(id, os.Args[3], os.Args[4], amountf)
		}
		fmt.Println(len(os.Args))
		if len(os.Args) < 5 {
			amount, _ := strconv.Atoi(os.Args[3])
			amountf := float64(amount)
			cli.UpdateExp(id, "", "", amountf)

		}
	}

}

func showUsgae() {
	fmt.Println("Show Usage")
}

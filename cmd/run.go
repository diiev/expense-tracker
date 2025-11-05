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
		if len(os.Args) < 3 {
			fmt.Println("Использование:\nupdate <id> [category] [description] [amount]\nupdate <id> <amount>\nupdate <id> <category> <description> <amount>\nupdate <id> <category> <amount>")
			return
		}

		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Некорректный ID")
			return
		}

		var category, description string
		var amount float64

		// Разбираем аргументы по количеству
		switch len(os.Args) {
		case 4:
			// update <id> <amount>
			amountInt, _ := strconv.Atoi(os.Args[3])
			amount = float64(amountInt)
		case 5:
			// update <id> <category> <amount>
			category = os.Args[3]
			amountInt, _ := strconv.Atoi(os.Args[4])
			amount = float64(amountInt)
		case 6:
			// update <id> <category> <description> <amount>
			category = os.Args[3]
			description = os.Args[4]
			amountInt, _ := strconv.Atoi(os.Args[5])
			amount = float64(amountInt)
		default:
			fmt.Println("Неверное количество аргументов.")
			fmt.Println("Пример: update 3 Еда Обед 500")
			return
		}

		if err := cli.UpdateExp(id, category, description, amount); err != nil {
			fmt.Println("Ошибка обновления:", err)
		}

	}

}

func showUsgae() {
	fmt.Println("Show Usage")
}

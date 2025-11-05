package cmd

import (
	"flag"
	"fmt"
	"os"

	"go.mod/internal/cli"
)

func Run() {
	if len(os.Args) < 2 {
		showUsage()
		return
	}

	command := os.Args[1]

	switch command {

	case "list":
		listCmd := flag.NewFlagSet("list", flag.ExitOnError)
		category := listCmd.String("category", "", "Фильтр по категории")
		_ = listCmd.Parse(os.Args[2:])
		cli.ShowExpense(*category)

	case "add":
		addCmd := flag.NewFlagSet("add", flag.ExitOnError)
		category := addCmd.String("category", "", "Категория расхода")
		desc := addCmd.String("desc", "", "Описание расхода")
		amount := addCmd.Float64("amount", 0, "Сумма расхода")
		_ = addCmd.Parse(os.Args[2:])

		if *category == "" || *desc == "" || *amount <= 0 {
			fmt.Println("Использование: add --category <категория> --desc <описание> --amount <сумма>")
			return
		}

		cli.AddExpense(*category, *desc, *amount)

	case "delete":
		deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
		id := deleteCmd.Int("id", 0, "ID записи для удаления")
		_ = deleteCmd.Parse(os.Args[2:])

		if *id <= 0 {
			fmt.Println("Использование: delete --id <номер>")
			return
		}

		cli.DeleteExpense(*id)

	case "summary":
		summaryCmd := flag.NewFlagSet("summary", flag.ExitOnError)
		month := summaryCmd.Int("month", 0, "Номер месяца (1-12)")
		_ = summaryCmd.Parse(os.Args[2:])

		cli.SummaryExp(*month)

	case "update":
		updateCmd := flag.NewFlagSet("update", flag.ExitOnError)
		id := updateCmd.Int("id", 0, "ID записи для обновления")
		category := updateCmd.String("category", "", "Категория расхода")
		desc := updateCmd.String("desc", "", "Описание расхода")
		amount := updateCmd.Float64("amount", 0, "Сумма расхода")
		_ = updateCmd.Parse(os.Args[2:])

		if *id <= 0 {
			fmt.Println("Использование: update --id <номер> [--category <категория>] [--desc <описание>] [--amount <сумма>]")
			return
		}

		if *category == "" && *desc == "" && *amount <= 0 {
			fmt.Println("Нужно указать хотя бы одно поле для обновления: --category, --desc, --amount")
			return
		}

		if err := cli.UpdateExp(*id, *category, *desc, *amount); err != nil {
			fmt.Println("Ошибка обновления:", err)
		}
	case "export":
		exportCmd := flag.NewFlagSet("export", flag.ExitOnError)
		filePath := exportCmd.String("file", "expenses.csv", "Имя файла для экспорта (по умолчанию expenses.csv)")
		category := exportCmd.String("category", "", "Фильтр по категории (необязательно)")
		_ = exportCmd.Parse(os.Args[2:])

		if err := cli.ExportToCSV(*filePath, *category); err != nil {
			fmt.Println("Ошибка экспорта:", err)
			return
		}
		fmt.Printf("✅ Данные успешно экспортированы в файл: %s\n", *filePath)
	case "set-budget":
		budgetCmd := flag.NewFlagSet("set-budget", flag.ExitOnError)
		month := budgetCmd.Int("month", 0, "Номер месяца (1–12)")
		amount := budgetCmd.Float64("amount", 0, "Сумма бюджета")
		_ = budgetCmd.Parse(os.Args[2:])

		if *month == 0 || *amount <= 0 {
			fmt.Println("Использование: set-budget --month <1-12> --amount <сумма>")
			return
		}

		if err := cli.SetBudget(*month, *amount); err != nil {
			fmt.Println("Ошибка установки бюджета:", err)
		}
	default:
		showUsage()
	}
}

func showUsage() {
	fmt.Println(`
Использование:
  list [--category <категория>]          Показать список расходов
  add --category <категория> --desc <описание> --amount <сумма>   Добавить расход
  delete --id <номер>                    Удалить запись
  summary [--month <1-12>]               Показать итоги за месяц
  update --id <номер> [--category <категория>] [--desc <описание>] [--amount <сумма>]   Обновить запись 
  set-budget --month <1-12> --amount <сумма>   Установить месячный бюджет
`)
}

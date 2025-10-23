package main

import "go.mod/internal/storage"

func main() {
	storage.LoadExpense(storage.FileName)
}

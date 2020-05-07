package presenters

import (
	"encoding/csv"
	"fmt"
	"io"

	"github.com/felipehfs/clean-api/entities"
)

// ExportBookToCSV export to csv the books
func ExportBookToCSV(books []entities.Book, output io.Writer) error {
	var sheet [][]string

	for _, book := range books {
		id := fmt.Sprintf("%v", book.ID)
		price := fmt.Sprintf("%v", book.Price)

		row := []string{id, book.Name, book.ISBN, price}
		sheet = append(sheet, row)
	}

	csvWriter := csv.NewWriter(output)
	return csvWriter.WriteAll(sheet)
}

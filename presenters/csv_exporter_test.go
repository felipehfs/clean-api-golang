package presenters_test

import (
	"os"

	"github.com/felipehfs/clean-api/entities"
	"github.com/felipehfs/clean-api/presenters"
)

func ExampleExportCSV() {
	books := []entities.Book{
		{ID: 1, Name: "Pequeno Príncipe", Price: 140.30, ISBN: "RERAIA-EIRURJGM-QQIW"},
	}
	presenters.ExportBookToCSV(books, os.Stdout)
	// output:
	// 1,Pequeno Príncipe,RERAIA-EIRURJGM-QQIW,140.3
}

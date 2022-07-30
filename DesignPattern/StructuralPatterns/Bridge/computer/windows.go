package computer

import (
	"Bridge/printer"
	"fmt"
)

type Windows struct {
	printer printer.Printer
}

func (w *Windows) Print() {
	fmt.Println("Print request for mac")
	w.printer.PrintFile()
}

func (w *Windows) SetPrinter(p printer.Printer) {
	w.printer = p
}

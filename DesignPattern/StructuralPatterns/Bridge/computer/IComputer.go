package computer

import "Bridge/printer"

type Computer interface {
	Print()
	SetPrinter(printer.Printer)
}

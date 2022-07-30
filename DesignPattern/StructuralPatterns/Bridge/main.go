package main

import (
	"Bridge/computer"
	"Bridge/printer"
	"fmt"
)

func main() {
	hpPrinter := &printer.Hp{}
	epsonPrinter := &printer.Epson{}

	macComputer := &computer.Mac{}

	macComputer.SetPrinter(hpPrinter)
	macComputer.Print()
	fmt.Println()

	macComputer.SetPrinter(epsonPrinter)
	macComputer.Print()
	fmt.Println()

	winComputer := &computer.Windows{}

	winComputer.SetPrinter(hpPrinter)
	winComputer.Print()
	fmt.Println()

	winComputer.SetPrinter(epsonPrinter)
	winComputer.Print()
	fmt.Println()
}

package main

import (
	"Adapter/client"
	"Adapter/computer"
)

func main() {
	cli := &client.Client{}
	mac := &computer.Mac{}

	cli.InsertLightningConnectorIntoComputer(mac)

	win := &computer.Windows{}
	winAdap := computer.NewWindowAdapter(win)

	cli.InsertLightningConnectorIntoComputer(winAdap)
}

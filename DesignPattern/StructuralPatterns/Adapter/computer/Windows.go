package computer

import "fmt"

type Windows struct {
}

func (w *Windows) insertIntoUSBPort() {
	fmt.Println("USB connector is plugged int windows machine.")
}

type WindowsAdapter struct {
	windowMachine *Windows
}

func NewWindowAdapter(windows *Windows) *WindowsAdapter {
	return &WindowsAdapter{
		windowMachine: windows,
	}
}

func (w *WindowsAdapter) InsertIntoLightningPort() {
	fmt.Println("Adapter converts Lightning signal to USB.")
	w.windowMachine.insertIntoUSBPort()
}

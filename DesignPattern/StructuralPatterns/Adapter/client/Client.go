package client

import (
	"Adapter/computer"
	"fmt"
)

type Client struct {
}

func (c *Client) InsertLightningConnectorIntoComputer(com computer.Computer) {
	fmt.Println("Client inserts Lighting connector into computer.")
	com.InsertIntoLightningPort()
}

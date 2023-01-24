package main

import (
	"log"
	"math/rand"
	"time"
)

type Seat int
type Bar chan Seat

func (bar Bar) ServeCustomerAtSeat(consumers chan int) {
	for c := range consumers {
		seatId := <-bar
		log.Print("++ customer#", c, " drinks at seat#", seatId)
		time.Sleep(time.Second * time.Duration(2+rand.Intn(6)))
		log.Print("-- customer#", c, " frees seat#", seatId)
		bar <- seatId // free seat and leave the bar
	}
}

func main() {
	rand.Seed(time.Now().UnixNano()) // before Go 1.20

	bar24x7 := make(Bar, 10)
	for seatId := 0; seatId < cap(bar24x7); seatId++ {
		bar24x7 <- Seat(seatId)
	}

	consumers := make(chan int)
	for i := 0; i < cap(bar24x7); i++ {
		go bar24x7.ServeCustomerAtSeat(consumers)
	}

	for customerId := 0; ; customerId++ {
		time.Sleep(time.Second)
		consumers <- customerId
	}
}

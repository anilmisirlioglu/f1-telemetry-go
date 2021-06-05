package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env/driver"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/telemetry"
)

func main() {
	client, err := telemetry.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Checo Sergio Perez Driver ID: %d\n", driver.SergioPerez)

	client.OnMotionPacket(func(packet *packets.PacketMotionData) {
		//fmt.Printf("%+v\n", packet)
	})
	client.OnEventPacket(func(packet *packets.PacketEventData) {
		fmt.Printf("Code: %s\n", packet.EventCodeString())
	})

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Printf("RecvCount: %d\n", client.Stats.RecvCount())
		fmt.Printf("ErrCount: %d\n", client.Stats.ErrCount())
		os.Exit(1)
	}()

	go func() {
		for range time.Tick(1 * time.Second) {
			time.Sleep(1 * time.Second)
			fmt.Printf("PPS: %d\n", client.Stats.PPS())
		}
	}()

	client.Run()
}

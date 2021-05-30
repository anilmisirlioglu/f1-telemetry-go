package main

import (
	"fmt"
	"log"

	"github.com/anilmisirlioglu/f1-telemetry/internal/packets"
	"github.com/anilmisirlioglu/f1-telemetry/pkg/env/driver"
	"github.com/anilmisirlioglu/f1-telemetry/pkg/telemetry"
)

func main() {
	client, err := telemetry.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(driver.SergioPerez)
	client.OnMotionPacket(func(packet *packets.PacketMotionData) {
		fmt.Printf("%+v\n", packet)
	})
	client.Run()
}

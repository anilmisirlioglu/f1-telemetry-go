package main

import (
	"fmt"
	packets2 "github.com/anilmisirlioglu/f1-telemetry/internal/packets"
	"log"

	"github.com/anilmisirlioglu/f1-telemetry/pkg/telemetry"
)

func main() {
	client, err := telemetry.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	client.OnMotionPacket(func(packet *packets2.PacketMotionData) {
		fmt.Printf("%+v\n", packet)
	})
	client.Run()
}

package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env/event"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/telemetry"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	// prometheus handler
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	client, err := telemetry.NewClient()
	if err != nil {
		log.Fatal(err)
	}

	// wait exit signal
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		log.Printf("Packet RecvCount: %d\n", client.Stats.RecvCount())
		log.Printf("Packet ErrCount: %d\n", client.Stats.ErrCount())
		os.Exit(1)
	}()

	// events
	client.OnEventPacket(func(packet *packets.PacketEventData) {
		switch packet.EventCodeString() {
		case event.SpeedTrapTriggered:
			trap := packet.EventDetails.(*packets.SpeedTrap)
			if trap.VehicleIdx == packet.Header.PlayerCarIndex {
				log.Printf("Speed Trap: %f\n", trap.Speed)
				speedTrapMetric.Set(float64(trap.Speed))
			}
		case event.FastestLap:
			fp := packet.EventDetails.(*packets.FastestLap)
			if fp.VehicleIdx == packet.Header.PlayerCarIndex {
				log.Printf("Fastest Lap: %f seconds", fp.LapTime)
				fastestLapMetric.Set(float64(fp.LapTime))
			}
		}
	})
	client.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		car := packet.CarTelemetryData[packet.Header.PlayerCarIndex]
		speedMetric.Set(float64(car.Speed))
		engineRPMMetric.Set(float64(car.EngineRPM))

		rlBrakeTempMetric.Set(float64(car.BrakesTemperature[0]))
		rrBrakeTempMetric.Set(float64(car.BrakesTemperature[1]))
		flBrakeTempMetric.Set(float64(car.BrakesTemperature[2]))
		frBrakeTempMetric.Set(float64(car.BrakesTemperature[3]))
	})
	client.OnLapPacket(func(packet *packets.PacketLapData) {
		lap := packet.LapData[packet.Header.PlayerCarIndex]
		lastLapTimeMetric.Set(float64(lap.LastLapTime))
	})

	log.Println("F1 telemetry client running")
	client.Run() // run F1 Telemetry Client
}

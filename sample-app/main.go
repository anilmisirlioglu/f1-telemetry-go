package main

import (
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"

	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/env/event"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/packets"
	"github.com/anilmisirlioglu/f1-telemetry-go/pkg/telemetry"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	speedTrapMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_speed_trap",
		Help: "Speed Trap Metric",
	})
	speedMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_car_speed",
		Help: "Car speed",
	}, []string{"speed", "engine_rpm", "throttle", "drs"})
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
		if packet.EventCodeString() == event.SpeedTrapTriggered {
			log.Printf("Speed Trap: %f\n", packet.EventDetails.SpeedTrap.Speed)
			speedTrapMetric.Set(float64(packet.EventDetails.SpeedTrap.Speed))
		}
	})
	client.OnCarTelemetryPacket(func(packet *packets.PacketCarTelemetryData) {
		car := packet.CarTelemetryData[0]
		speedMetric.WithLabelValues(
			strconv.Itoa(int(car.Speed)),
			strconv.Itoa(int(car.EngineRPM)),
			strconv.Itoa(int(car.Throttle)),
			strconv.FormatBool(car.DRS != 0),
		)
	})
	log.Println("F1 telemetry client running")
	client.Run() // run F1 Telemetry Client
}

package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	speedTrapMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_speed_trap",
		Help: "Speed Trap",
	})
	speedMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_car_speed",
		Help: "Car Speed",
	})
	engineRPMMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_car_engine_rpm",
		Help: "Car Engine RPM",
	})
	fastestLapMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_fastest_lap",
		Help: "Fastest Lap",
	})
	lastLapTimeMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_last_lap",
		Help: "Last Lap",
	})
	// Brake Temp Metrics
	rlBrakeTempMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_brake_temp_rl",
		Help: "Rear Left Tyre Brake Temperature",
	})
	rrBrakeTempMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_brake_temp_rr",
		Help: "Rear Right Tyre Brake Temperature",
	})
	flBrakeTempMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_brake_temp_fl",
		Help: "Front Left Tyre Brake Temperature",
	})
	frBrakeTempMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_brake_temp_fr",
		Help: "Front Right Tyre Brake Temperature",
	})
)

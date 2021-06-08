package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var (
	// Speed Metric
	speedTrapMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_speed_trap",
		Help: "Speed Trap",
	})
	speedMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_car_speed",
		Help: "Car Speed",
	})

	// Engine Metrics
	engineRPMMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_car_engine_rpm",
		Help: "Car Engine RPM",
	})
	engineDamageMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_engine_damage",
		Help: "Car Engine Damage",
	})

	// Lap Time Metrics
	fastestLapMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_fastest_lap",
		Help: "Fastest Lap Time",
	})
	lastLapTimeMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_last_lap",
		Help: "Last Lap Time",
	})

	// Tyre Metric
	tyresAgeLapsMetric = promauto.NewGauge(prometheus.GaugeOpts{
		Name: "f1_telemetry_tyres_age_lap",
		Help: "Tyres Age Lap",
	})
	tyreWearMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_tyre_wear",
		Help: "Tyres Wear",
	}, []string{"tyre"})

	// Brake Temp Metrics
	brakesTempMetric = promauto.NewGaugeVec(prometheus.GaugeOpts{
		Name: "f1_telemetry_brake_temp",
		Help: "Brake Temperatures",
	}, []string{"brake"})
)

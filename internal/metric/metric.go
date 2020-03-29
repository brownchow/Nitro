package metric

import (
	"time"

	"github.com/influxdata/influxdb-client-go"
)

type Client interface {
	Add(metric Metric)
	Send(metrics ...influxdb.Metric)
	Ticker(duration time.Duration)
}

type Metric interface {
	Metric() influxdb.Metric
}

type RowMetric struct {
	Name   string
	Tags   Tags
	Fields Fields
	Time   time.Time
}

func (rm *RowMetric) Metric() influxdb.Metric {
	return influxdb.NewRowMetric(rm.Fields, rm.Name, rm.Tags, rm.Time)
}

type CounterMetric struct {
	RowMetric
	Counter int
}

func (cm *CounterMetric) Inc() {
	cm.Counter++
	cm.Time = time.Now()
}

func (cm *CounterMetric) Metric() influxdb.Metric {
	return influxdb.NewRowMetric(Fields{"counter": cm.Counter}, cm.Name, cm.Tags, cm.Time)
}

type GaugeMetric struct {
	RowMetric
	Gauge int
}

func (gm *GaugeMetric) Set(gauge int) {
	gm.Gauge = gauge
	gm.Time = time.Now()
}

func (gm *GaugeMetric) Metric() influxdb.Metric {
	return influxdb.NewRowMetric(Fields{"gauge": gm.Gauge}, gm.Name, gm.Tags, gm.Time)
}
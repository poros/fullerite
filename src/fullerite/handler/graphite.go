package handler

import (
	"fullerite/metric"
	"log"
)

// Graphite type
type Graphite struct {
	BaseHandler
	server string
}

// NewGraphite returns a new Graphite handler.
func NewGraphite() *Graphite {
	g := new(Graphite)
	g.name = "SignalFx"
	g.interval = DefaultInterval
	g.maxBufferSize = DefaultBufferSize
	g.channel = make(chan metric.Metric)

	return g
}

// Run sends metrics in the channel to the graphite server.
func (g Graphite) Run() {
	// TODO: check interval and queue size and metrics.
	for metric := range g.channel {
		// TODO: Actually send to graphite server
		log.Println("Sending metric to Graphite:", metric)
	}

}

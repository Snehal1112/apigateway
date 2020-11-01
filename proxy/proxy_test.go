package proxy

import (
	"encoding/json"
	"log"
	"testing"
)

func setup() *Proxy {
	return NewProxy()
}

func TestNewProxy(t *testing.T) {
	var targets []Target
	targets = append(targets, *NewTarget("http://service1:8080/"), *NewTarget("http://service2:8080/"))

	Upstreams := NewUpstreams(WithBalancing("roundrobin"), WithTargets(targets))
	Proxy := NewProxy(
		WithPreserveHost(false),
		WithListenPath("https://www.google.com"),
		WithStripPath(false),
		WithAppendPath(true),
		WithMethods([]string{"GET", "POST"}),
		WithUpstreams(Upstreams),
	)

	b, err := json.MarshalIndent(Proxy, "", "    ")
	log.Println(err)
	log.Println(string(b))
}

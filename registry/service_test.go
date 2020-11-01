package registry

import (
	"encoding/json"
	"github.com/snehal1112/gateway/proxy"
	"log"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	var targets []proxy.Target
	targets = append(targets, *proxy.NewTarget("http://service1:8080/"), *proxy.NewTarget("http://service2:8080/"))

	Upstreams := proxy.NewUpstreams(proxy.WithBalancing("roundrobin"), proxy.WithTargets(targets))

	re := NewService("SD", true)
	re.Proxy = re.NewProxy(		proxy.WithPreserveHost(false),
		proxy.WithListenPath("https://www.google.com"),
		proxy.WithStripPath(false),
		proxy.WithAppendPath(true),
		proxy.WithMethods([]string{"GET", "POST"}),
		proxy.WithUpstreams(Upstreams),
	)
	b, err := json.MarshalIndent(re, "", "    ")
	log.Println("err:-", err)
	log.Println("b:-", string(b))
}

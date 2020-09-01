package registry

import (
	"encoding/json"
	"github.com/snehal1112/gateway/proxy"
	"log"
	"testing"
)

func TestNewRegistry(t *testing.T) {
	re := NewService("SD", true, proxy.NewProxy())
	log.Println(re)

	b, err := json.MarshalIndent(re, "", "    ")
	log.Println("err:-", err)
	log.Println("b:-", string(b))
}

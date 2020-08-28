package bootstrap

import "log"

type Server struct {
	configFile string
}

func NewServer(options ...Option)  {
	ser := &Server{}

	for _, option := range options {
		option(ser)
	}

	log.Println(ser.configFile)
}
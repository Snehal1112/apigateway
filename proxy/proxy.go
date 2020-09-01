package proxy

type Target struct {
	Target string `json:"target"`
}

type Upstreams struct {
	Balancing string   `bson:"balancing"`
	Targets   []Target `bson:"targets"`
}

type Proxy struct {
	PreserveHost bool      `bson:"preserve_host"`
	ListenPath   string    `bson:"listen_path"`
	StripPath    bool      `bson:"strip_path"`
	AppendPath   bool      `bson:"append_path"`
	Methods      []string  `bson:"methods"`
	Upstreams    Upstreams `bson:"upstreams"`
}

func NewProxy(options ...OptionProxy) *Proxy {
	return newProxy(options...)
}

func newProxy(options ...OptionProxy) *Proxy {
	proxy := &Proxy{}
	for _, option := range options {
		option(proxy)
	}

	return proxy
}

func NewUpstreams(options ...OptionUpstreams) *Upstreams {
	upstreams := &Upstreams{}
	for _, option := range options {
		option(upstreams)
	}
	return upstreams
}

func NewTarget(target string) *Target {
	return &Target{Target: target}
}

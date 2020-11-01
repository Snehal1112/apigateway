package proxy

type OptionProxy func(a *Proxy)

func WithPreserveHost(preserveHost bool) OptionProxy {
	return func(a *Proxy) {
		a.PreserveHost = preserveHost
	}
}

func WithListenPath(listenPath string) OptionProxy {
	return func(a *Proxy) {
		a.ListenPath = listenPath
	}
}

func WithStripPath(stripPath bool) OptionProxy {
	return func(a *Proxy) {
		a.StripPath = stripPath
	}
}

func WithAppendPath(appendPath bool) OptionProxy {
	return func(a *Proxy) {
		a.AppendPath = appendPath
	}
}

func WithMethods(methods []string) OptionProxy {
	return func(a *Proxy) {
		a.Methods = methods
	}
}

func WithUpstreams(upstreams *Upstreams) OptionProxy {
	return func(a *Proxy) {
		a.Upstreams = *upstreams
	}
}

type OptionUpstreams func(u *Upstreams)

func WithBalancing(balancing string) OptionUpstreams {
	return func(u *Upstreams) {
		u.Balancing = balancing
	}
}

func WithTargets(targets []Target) OptionUpstreams {
	return func(u *Upstreams) {
		u.Targets = targets
	}
}

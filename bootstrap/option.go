package bootstrap

type Option func(s *Server)

func ConfigFile(file string) Option {
	return func(s *Server) {
		s.configFile = file
	}
}
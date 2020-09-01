package cmd

import "os"

func getEnv(name string, def string) string {
	v := os.Getenv(name)
	if v == "" {
		return def
	}

	return v
}

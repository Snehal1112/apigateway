/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"context"
	"fmt"
	"github.com/snehal1112/gateway/bootstrap"
	"github.com/snehal1112/gateway/config"
	"github.com/snehal1112/gateway/server"
	"os"

	"github.com/spf13/cobra"
)

const (
	defaultListenAddr = "127.0.0.1:8773"
	uriBasePath       = "/api/v1"
	defaultDBURI = "mongodb://0.0.0.0:27017/?retryWrites=false"
	defaultDatabase = "gateway"
)

var bootstrapConfig = &bootstrap.Config{}

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Api gateway service which act as a gateway of all micro service.",
	Run: func(cmd *cobra.Command, args []string) {
		if err := serve(cmd, args); err != nil {
			fmt.Printf("Error: %v \n\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	cfg := bootstrapConfig
	serveCmd.Flags().StringVar(&cfg.Listen, "listen", getEnv("GATEWAY_LISTEN", defaultListenAddr), fmt.Sprintf("TCP listen address (default \"%s\").", "8773"))
	serveCmd.Flags().StringVar(&cfg.URIBasePath, "api_base", getEnv("GATEWAY_BASE_API", uriBasePath), "uri base path for an api gateway.")
	serveCmd.Flags().StringVar(&cfg.BackendURL, "backend_url", getEnv("GATEWAY_BACKEND_URL", defaultDBURI), "uri base path for an api gateway backend.")
	serveCmd.Flags().StringVar(&cfg.DatabaseName, "database_name", getEnv("GATEWAY_DATABASE", defaultDatabase), "database which used be the api gateway")
	serveCmd.Flags().Bool("log-timestamp", true, "Prefix each log line with timestamp")
	serveCmd.Flags().String("log-level", "info", "Log level (one of panic, fatal, error, warn, info or debug)")
}

func serve(cmd *cobra.Command, args []string) error {
	ctx := context.Background()
	logTimestamp, _ := cmd.Flags().GetBool("log-timestamp")
	logLevel, _ := cmd.Flags().GetString("log-level")

	logger, err := newLogger(!logTimestamp, logLevel)
	if err != nil {
		return fmt.Errorf("failed to create logger: %v", err)
	}
	logger.Infoln("serve start")

	bs, err := bootstrap.Boot(ctx, bootstrapConfig, &config.Config{
		Logger: logger,
	})

	if err != nil {
		return err
	}

	srv, err := server.NewServer(
		server.WithLogger(logger),
		server.WithConfig(&server.Config{
			Config:   bs.Config(),
			Services: bs.Manager().Services(),
			BasePath: bootstrapConfig.URIBasePath,
		}),
		server.WithListener(bs.Config().ListenAddr),
	)
	if err != nil {
		return fmt.Errorf("failed to create server: %v", err)
	}

	srv.Serve(ctx)

	return nil
}

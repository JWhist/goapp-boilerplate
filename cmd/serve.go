package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/JWhist/jw-goapp/config"
	"github.com/JWhist/jw-goapp/server"
	"github.com/spf13/cobra"
)

var configPath string

var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start server",
	Long:  `Starts a generic API server at the configuration specified in the yaml file.`,
	Run: func(cmd *cobra.Command, args []string) {
		serve()
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)
	serveCmd.PersistentFlags().StringVarP(&configPath, "config", "c", "config/config.yaml", "Path to config file")
}

func serve() {
	fmt.Printf("serve called using config at %v\n", configPath)
	config := config.NewConfig(configPath)
	fmt.Printf("Using port %v\n", config.Port)
	server := server.NewServer(config)
	fmt.Println("Starting server")

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("Received %v, shutting down\n", sig)
		server.Close()
	}()

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
	defer server.Close()
}

package main

import (
	"context"
	"fmt"
	"hw/sources"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var configFile string

func printVersion() {
	fmt.Println("v.1")
}

func init() {
	pflag.StringVar(&configFile, "config", "", "Path to configuration file")
}

func main() {
	pflag.Parse()
	if pflag.Arg(0) == "version" {
		printVersion()
		return
	}
	if configFile == "" {
		fmt.Println("Please set: '--config <Path to configuration file>'")
		return
	}
	viper.SetConfigType("yaml")
	file, err := os.Open(configFile)
	if err != nil {
		fmt.Println(err)
		return
	}
	viper.ReadConfig(file)
	httpConfig := sources.HTTPConfig{}
	err = viper.Unmarshal(&httpConfig) // TODO
	if err != nil {
		fmt.Printf("unable to decode into struct, %v\n", err)
		return
	}

	httpServer := sources.NewHTTPServer(
		httpConfig.Host,
		httpConfig.Port,
	)

	ctx, stop := signal.NotifyContext(context.Background(),
		syscall.SIGINT, syscall.SIGTERM, syscall.SIGHUP, syscall.SIGTSTP)
	defer stop()
	wg := sync.WaitGroup{}

	var once sync.Once
	// GRASEFULL: httpServer.Stop
	wg.Add(1)
	go func() {
		defer wg.Done()
		<-ctx.Done()
		if err := httpServer.Stop(ctx); err != nil {
			fmt.Println(err.Error())
		}
	}()

	// httpServer.Start
	wg.Add(1)
	go func() {
		defer wg.Done()
		if err := httpServer.Start(); err != nil {
			fmt.Println("failed to start HTTP server: " + err.Error())
			once.Do(stop)
		}
	}()

	fmt.Println("Http was running...")
	<-ctx.Done()
	fmt.Println("Complex Shutting down was done gracefully by signal.")
	wg.Wait()
}

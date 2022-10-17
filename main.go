package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/Ning-Qing/temple/config"
	"github.com/Ning-Qing/temple/global"
	"github.com/Ning-Qing/temple/router"
	"github.com/spf13/cobra"
)

var (
	application string
	version     string
)

var cmd = &cobra.Command{
	Use:     application,
	Short:   "",
	Version: version,
	Run:     run,
}

func init() {
	global.Application = application
	global.Version = version
	initFlage()
}

func main() {
	if err := cmd.Execute(); err != nil {
		log.Fatalf("start server faild: %s", err.Error())
	}
}

func initFlage() {
	cmd.PersistentFlags().StringVarP(&global.ConfigPath, "config", "c", "config.yaml", "the path to the configuration file")
}

func run(cmd *cobra.Command, args []string) {
	global.GlobalSettings = config.InitConfig(global.ConfigPath)
	r := router.InitRouter()
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", global.GlobalSettings.GetServerHost(), global.GlobalSettings.GetServerPort()),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	exit(srv, quit)
}

func exit(srv *http.Server, quit chan os.Signal) {
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}

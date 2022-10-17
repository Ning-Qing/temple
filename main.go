package main

import (
	"fmt"
	"os"

	"github.com/Ning-Qing/temple/config"
	"github.com/Ning-Qing/temple/router"
	"github.com/spf13/cobra"
)

var cmd = &cobra.Command{
	Use:     "temple",
	Short:   "",
	Version: Version,
	Run:     run,
}

var (
	Version    string
	ConfigPath string
)

func init() {
	cmd.PersistentFlags().StringVarP(&ConfigPath, "config", "c", "config.yaml", "the path to the configuration file")
}

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func run(cmd *cobra.Command, args []string) {
	config.InitConfig(ConfigPath)
	r := router.InitRouter()
	r.Run(fmt.Sprintf("%s:%s", config.GlobalSettings.Server.Host, config.GlobalSettings.Server.Port))
}

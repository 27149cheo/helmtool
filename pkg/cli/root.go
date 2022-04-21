package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/27149cheo/helmtool/pkg/tool/log"
)

const (
	namespace  = "namespace"
	kubeConfig = "kubeconfig"
)

var rootCmd = &cobra.Command{
	Use:   "helmtool",
	Short: "set of helm tools",
	Long:  `set of helm tools.`,
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringP(namespace, "n", "default", "namespace of the release")
	rootCmd.PersistentFlags().StringP(kubeConfig, "", "", "path to the kubeconfig file")

	_ = viper.BindPFlag(namespace, rootCmd.PersistentFlags().Lookup(namespace))
	_ = viper.BindPFlag(kubeConfig, rootCmd.PersistentFlags().Lookup(kubeConfig))
}

func initConfig() {
	viper.AutomaticEnv()

	log.Init(&log.Config{
		Level:    "debug",
		NoCaller: true,
	})
}

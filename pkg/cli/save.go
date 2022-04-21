package cli

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/apimachinery/pkg/runtime"
	clientgoscheme "k8s.io/client-go/kubernetes/scheme"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/clientcmd"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/27149cheo/helmtool/pkg/cli/save"
	"github.com/27149cheo/helmtool/pkg/tool/log"
)

const (
	destination = "destination"
)

const saveDesc = `
This command reads the latest revision of a release from secret 
and creates an archived chart in the given directory.
`

var cl client.Client

func init() {
	rootCmd.AddCommand(saveCmd)

	saveCmd.PersistentFlags().StringP(destination, "o", "/tmp", "path to save the chart tarball")
	_ = viper.BindPFlag(destination, saveCmd.PersistentFlags().Lookup(destination))
}

var saveCmd = &cobra.Command{
	Use:   "save RELEASE_NAME [flags]",
	Short: "save a chart from release secret",
	Long:  saveDesc,
	Args:  cobra.ExactArgs(1),
	PreRunE: func(cmd *cobra.Command, args []string) error {
		return preRun()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if err := run(args); err != nil {
			log.Fatal(err)
		}
	},
}

func run(args []string) error {
	return save.Save(
		viper.GetString(namespace),
		args[0],
		viper.GetString(destination),
		cl)
}

func preRun() error {
	scheme := runtime.NewScheme()
	_ = clientgoscheme.AddToScheme(scheme)

	restConfig, err := getConfig()
	if err != nil {
		return err
	}

	cl, err = client.New(restConfig, client.Options{
		Scheme: scheme,
	})

	return err
}

func getConfig() (*rest.Config, error) {
	if viper.GetString(kubeConfig) != "" {
		return clientcmd.NewNonInteractiveDeferredLoadingClientConfig(
			&clientcmd.ClientConfigLoadingRules{ExplicitPath: viper.GetString(kubeConfig)},
			&clientcmd.ConfigOverrides{},
		).ClientConfig()
	}

	return ctrl.GetConfig()
}

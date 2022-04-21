package save

import (
	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/chartutil"
	"helm.sh/helm/v3/pkg/storage"
	"helm.sh/helm/v3/pkg/storage/driver"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/yaml"

	"github.com/27149cheo/helmtool/pkg/tool/log"
)

// Save reads the latest revision of a release, and creates an archived chart to the given directory.
func Save(namespace, releaseName, dir string, cl client.Client) error {
	log.Infof("Processing release %s in namespace %s", releaseName, namespace)
	d := driver.NewSecrets(NewSecretReader(namespace, cl))
	store := storage.Init(d)

	release, err := store.Last(releaseName)
	if err != nil {
		return err
	}

	ch := release.Chart
	values, err := yaml.Marshal(ch.Values)
	if err != nil {
		return err
	}

	// Raw is empty for a chart loaded from a release secret, but the values.yaml is required in chartutil.Save,
	// so we populate it manually.
	ch.Raw = append(ch.Raw, &chart.File{Name: chartutil.ValuesfileName, Data: values})

	log.Infof("Saving to %s", dir)
	res, err := chartutil.Save(ch, dir)
	if err != nil {
		return err
	}

	log.Infof("An archived chart is created in %s", res)

	return nil
}

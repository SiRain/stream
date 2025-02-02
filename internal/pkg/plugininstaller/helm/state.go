package helm

import (
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller"
	"github.com/devstream-io/devstream/internal/pkg/plugininstaller/common"
	"github.com/devstream-io/devstream/internal/pkg/statemanager"
	"github.com/devstream-io/devstream/pkg/util/helm"
)

// GetPluginAllState will get the State of k8s Deployment, DaemonSet and StatefulSet resources
func GetPluginAllState(options plugininstaller.RawOptions) (statemanager.ResourceState, error) {
	opts, err := NewOptions(options)
	if err != nil {
		return nil, err
	}

	anFilter := map[string]string{
		helm.GetAnnotationName(): opts.GetReleaseName(),
	}
	labelFilter := map[string]string{}
	return common.GetPluginAllK8sState(opts.GetNamespace(), anFilter, labelFilter)
}

package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// KindKubeControllersConfiguration const
	KindKubeControllersConfiguration     = "KubeControllersConfiguration"
	// KindKubeControllersConfigurationList const
	KindKubeControllersConfigurationList = "KubeControllersConfigurationList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeControllersConfiguration is the Schema for the kubecontrollersconfigurations API
type KubeControllersConfiguration struct {
	metav1.TypeMeta   `json:",inline"`
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   KubeControllersConfigurationSpec   `json:"spec,omitempty"`
	// +optional
	Status KubeControllersConfigurationStatus `json:"status,omitempty"`
}

// KubeControllersConfigurationSpec defines the desired state of KubeControllersConfiguration
type KubeControllersConfigurationSpec struct {
	// LogSeverity is the log severity above which logs are sent to the stdout. [Default: Info]
	LogSeverity string `json:"logSeverity,omitempty" validate:"omitempty"`

	// HealthChecks enables or disables support for health checks [Default: Enabled]
	HealthChecks string `json:"healthChecks,omitempty" validate:"omitempty,oneof=Enabled Disabled"`

	// Controllers enables and configures individual Kubernetes controllers
	Controllers ControllersConfig `json:"controllers"`
}

// ControllersConfig enables and configures individual Kubernetes controllers
type ControllersConfig struct {
	// Node enables and configures the node controller. Enabled by default, set to nil to disable.
	Node *NodeControllerConfig `json:"node,omitempty"`

	// Namespace enables and configures the namespace controller. Enabled by default, set to nil to disable.
	Namespace *NamespaceControllerConfig `json:"namespace,omitempty"`

	// Multus enables and configures the multus controller. Enabled by default, set to nil to disable.
	Multus *MultusControllerConfig `json:"multus,omitempty"`

	// ConfigMap enables and configures the ConfigMap controller. Enabled by default, set to nil to disable.
	ConfigMap *ConfigMapControllerConfig `json:"configMap,omitempty"`
}

// NodeControllerConfig configures the node controller, which automatically cleans up configuration
// for nodes that no longer exist.
type NodeControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the fsc datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// NamespaceControllerConfig configures the namespace controller, which identifies the active namespaces in the cluster
type NamespaceControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the fsc datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// MultusControllerConfig configures the multus controller, which identifies the configured network attachements in the cluster
type MultusControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the fsc datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// ConfigMapControllerConfig configures the configMap controller, which identifies the configured sriov config in the cluster
type ConfigMapControllerConfig struct {
	// ReconcilerPeriod is the period to perform reconciliation with the fsc datastore. [Default: 5m]
	ReconcilerPeriod *metav1.Duration `json:"reconcilerPeriod,omitempty" validate:"omitempty"`
}

// KubeControllersConfigurationStatus represents the status of the configuration. It's useful for admins to
// be able to see the actual config that was applied, which can be modified by environment variables on the
// kube-controllers process.
type KubeControllersConfigurationStatus struct {
	// RunningConfig contains the effective config that is running in the kube-controllers pod, after
	// merging the API resource with any environment variables.
	RunningConfig KubeControllersConfigurationSpec `json:"runningConfig,omitempty"`

	// EnvironmentVars contains the environment variables on the kube-controllers that influenced
	// the RunningConfig.
	EnvironmentVars map[string]string `json:"environmentVars,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// KubeControllersConfigurationList contains a list of KubeControllersConfiguration
type KubeControllersConfigurationList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []KubeControllersConfiguration `json:"items"`
}

// NewKubeControllersConfiguration creates a new (zeroed) KubeControllersConfiguration struct with
// the TypeMetadata initialized to the current version.
func NewKubeControllersConfiguration() *KubeControllersConfiguration {
	return &KubeControllersConfiguration{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindKubeControllersConfiguration,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}

// NewKubeControllersConfigurationList creates a new (zeroed) KubeControllersConfigurationList struct with the TypeMetadata
// initialized to the current version.
func NewKubeControllersConfigurationList() *KubeControllersConfigurationList {
	return &KubeControllersConfigurationList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindKubeControllersConfigurationList,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}
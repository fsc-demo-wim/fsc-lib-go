package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// KindWorkLoad const
	KindWorkLoad = "WorkLoad"
	// KindWorkLoadList const
	KindWorkLoadList = "WorkLoadList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// WorkLoad is the Schema for the workload API
type WorkLoad struct {
	metav1.TypeMeta `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec WorkLoadSpec `json:"spec,omitempty"`

	Status WorkLoadStatus `json:"status,omitempty"`
}

// WorkLoadSpec defines the desired state of WorkLoad
type WorkLoadSpec struct {
	// RouteTarget identifies the route target of the workload
	RouteTarget string `json:"routeTarget,omitempty"`

	// RouteDistinguisher identifies the route distinguisher of the workload
	RouteDistinguisher string `json:"routeDistinguisher,omitempty"`

	// Vlans defines the vlans attached to the workload
	Vlans []*Vlan `json:"vlans,omitempty"`
}

// WorkLoadStatus represents the status of the workload configuration
// and the associated endpoints and provisioning status
type WorkLoadStatus struct {
	// RunningConfig contains the effective config that is running, after
	// merging the API resource with any nodal interface/vlan status
	RunningConfig WorkLoadSpec `json:"runningConfig,omitempty"`

	// Devices and its related interface on which this workload is applied
	Devices []*Device `json:"devices,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkLoadList contains a list of workloads
type WorkLoadList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []WorkLoad `json:"items"`
}

// NewWorkLoad creates a new (zeroed) Workload struct with
// the TypeMetadata initialized to the current version.
func NewWorkLoad() *WorkLoad {
	return &WorkLoad{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindWorkLoad,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}

// NewWorkLoadList creates a new (zeroed) WorkloadList struct with the TypeMetadata
// initialized to the current version.
func NewWorkLoadList() *WorkLoadList {
	return &WorkLoadList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindWorkLoadList,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}

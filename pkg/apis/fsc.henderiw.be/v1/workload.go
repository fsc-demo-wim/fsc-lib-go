package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// KindWorkload const
	KindWorkload    = "Workload"
	// KindWorkloadList const
	KindWorkloadList = "WorkloadList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// Workload is the Schema for the workload API
type Workload struct {
	metav1.TypeMeta   `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   WorkloadSpec   `json:"spec,omitempty"`

	Status WorkloadStatus `json:"status,omitempty"`
}

// WorkloadSpec defines the desired state of Workload
type WorkloadSpec struct {
	// RouteTarget identifies the route target of the workload
	RouteTarget string `json:"routeTarget,omitempty"`

	// RouteDistinguisher identifies the route distinguisher of the workload
	RouteDistinguisher string `json:"routeDistinguisher,omitempty"`	

	// Vlans defines the vlans attached to the workload
	Vlans []*Vlan `json:"vlans,omitempty"`
}

// WorkloadStatus represents the status of the workload configuration
// and the associated endpoints and provisioning status
type WorkloadStatus struct {
	// RunningConfig contains the effective config that is running, after
	// merging the API resource with any nodal interface/vlan status
	RunningConfig WorkloadSpec `json:"runningConfig,omitempty"`

	// Devices and its related interface on which this workload is applied
	Devices []*Device `json:"devices,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// WorkloadList contains a list of workloads
type WorkloadList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Workload`json:"items"`
}

// NewWorkload creates a new (zeroed) Workload struct with
// the TypeMetadata initialized to the current version.
func NewWorkload() *Workload{
	return &Workload{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindWorkload,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}

// NewWorkloadList creates a new (zeroed) WorkloadList struct with the TypeMetadata
// initialized to the current version.
func NewWorkloadList() *WorkloadList {
	return &WorkloadList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindWorkloadList,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}
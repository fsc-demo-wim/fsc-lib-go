package v1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

const (
	// KindNodeTopology const
	KindNodeTopology     = "NodeTopology"
	// KindNodeTopologyList const
	KindNodeTopologyList = "NodeTopologyList"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true

// NodeTopology is the Schema for the node topology API
type NodeTopology struct {
	metav1.TypeMeta   `json:",inline"`

	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   NodeTopologySpec   `json:"spec,omitempty"`
}

// NodeTopologySpec defines the desired state of NodeTopology
type NodeTopologySpec struct {
	// Devices identifies the devices connected to the node
	Devices []*Device `json:"devices,omitempty"`
}

// Device defines the Device attibutes
type Device struct {
	// Name defines the device Name
	Name string `json:"name,omitempty"`

	// Kind defines the kind of device, like wbx, 7220, etc
	Kind string `json:"kind,omitempty"`

	// DeviceIndentifier defines the id of the device
	DeviceIdentifier string `json:"deviceIdentifier,omitempty"`

	// DeviceIndentifier defines the type of the device identifier
	DeviceIdentifierType string `json:"deviceIdentifierType,omitempty"`

	// Endpoints defines the endpoints attached to the device
	Endpoints []*Endpoint `json:"endpoints,omitempty"`
}

// Endpoint defines the Endpoint attributes
type Endpoint struct {
	// Name defines the endpoint Name
	Name string `json:"name,omitempty"`

	// InterfaceIdentifier defines the identifier of the interface
	InterfaceIdentifier string `json:"interfaceIdentifier,omitempty"`

	// InterfaceIdentifierType defines the type of the interface identifier
	InterfaceIdentifierType string `json:"interfaceIdentifierType,omitempty"`

	// MTU defines the endpoint MTU
	MTU int `json:"mtu,omitempty"`

	// LAG identifies if the interface is part of a LAG
	LAG bool `json:"lag,omitempty"`

	// Vlans defines the vlans attached to the endpoint
	Vlans []*Vlan `json:"vlans,omitempty"`
}

// Vlan defines the vlan attributes
type Vlan struct {
	// ID defines the vlan identifier
	ID int `json:"id,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// NodeTopologyList contains a list of NodeTopology
type NodeTopologyList struct {
	metav1.TypeMeta `json:",inline"`
	// +optional
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []NodeTopology `json:"items"`
}

// NewNodeTopology creates a new (zeroed) NodeTopology struct with
// the TypeMetadata initialized to the current version.
func NewNodeTopology() *NodeTopology {
	return &NodeTopology{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindNodeTopology,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}

// NewNodeTopologyList creates a new (zeroed) NodeTopologyList struct with the TypeMetadata
// initialized to the current version.
func NewNodeTopologyList() *NodeTopologyList {
	return &NodeTopologyList{
		TypeMeta: metav1.TypeMeta{
			Kind:       KindNodeTopologyList,
			APIVersion: SchemeGroupVersion.String(),
		},
	}
}
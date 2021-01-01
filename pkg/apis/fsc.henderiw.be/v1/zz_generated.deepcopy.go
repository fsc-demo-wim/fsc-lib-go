// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConfigMapControllerConfig) DeepCopyInto(out *ConfigMapControllerConfig) {
	*out = *in
	if in.ReconcilerPeriod != nil {
		in, out := &in.ReconcilerPeriod, &out.ReconcilerPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConfigMapControllerConfig.
func (in *ConfigMapControllerConfig) DeepCopy() *ConfigMapControllerConfig {
	if in == nil {
		return nil
	}
	out := new(ConfigMapControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllersConfig) DeepCopyInto(out *ControllersConfig) {
	*out = *in
	if in.Node != nil {
		in, out := &in.Node, &out.Node
		*out = new(NodeControllerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = new(NamespaceControllerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Multus != nil {
		in, out := &in.Multus, &out.Multus
		*out = new(MultusControllerConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.ConfigMap != nil {
		in, out := &in.ConfigMap, &out.ConfigMap
		*out = new(ConfigMapControllerConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllersConfig.
func (in *ControllersConfig) DeepCopy() *ControllersConfig {
	if in == nil {
		return nil
	}
	out := new(ControllersConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Device) DeepCopyInto(out *Device) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]*Endpoint, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Endpoint)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Device.
func (in *Device) DeepCopy() *Device {
	if in == nil {
		return nil
	}
	out := new(Device)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Endpoint) DeepCopyInto(out *Endpoint) {
	*out = *in
	if in.Vlans != nil {
		in, out := &in.Vlans, &out.Vlans
		*out = make([]*Vlan, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Vlan)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Endpoint.
func (in *Endpoint) DeepCopy() *Endpoint {
	if in == nil {
		return nil
	}
	out := new(Endpoint)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllersConfiguration) DeepCopyInto(out *KubeControllersConfiguration) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllersConfiguration.
func (in *KubeControllersConfiguration) DeepCopy() *KubeControllersConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeControllersConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllersConfiguration) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllersConfigurationList) DeepCopyInto(out *KubeControllersConfigurationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KubeControllersConfiguration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllersConfigurationList.
func (in *KubeControllersConfigurationList) DeepCopy() *KubeControllersConfigurationList {
	if in == nil {
		return nil
	}
	out := new(KubeControllersConfigurationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KubeControllersConfigurationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllersConfigurationSpec) DeepCopyInto(out *KubeControllersConfigurationSpec) {
	*out = *in
	if in.Controllers != nil {
		in, out := &in.Controllers, &out.Controllers
		*out = new(ControllersConfig)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllersConfigurationSpec.
func (in *KubeControllersConfigurationSpec) DeepCopy() *KubeControllersConfigurationSpec {
	if in == nil {
		return nil
	}
	out := new(KubeControllersConfigurationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeControllersConfigurationStatus) DeepCopyInto(out *KubeControllersConfigurationStatus) {
	*out = *in
	in.RunningConfig.DeepCopyInto(&out.RunningConfig)
	if in.EnvironmentVars != nil {
		in, out := &in.EnvironmentVars, &out.EnvironmentVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeControllersConfigurationStatus.
func (in *KubeControllersConfigurationStatus) DeepCopy() *KubeControllersConfigurationStatus {
	if in == nil {
		return nil
	}
	out := new(KubeControllersConfigurationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultusControllerConfig) DeepCopyInto(out *MultusControllerConfig) {
	*out = *in
	if in.ReconcilerPeriod != nil {
		in, out := &in.ReconcilerPeriod, &out.ReconcilerPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultusControllerConfig.
func (in *MultusControllerConfig) DeepCopy() *MultusControllerConfig {
	if in == nil {
		return nil
	}
	out := new(MultusControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceControllerConfig) DeepCopyInto(out *NamespaceControllerConfig) {
	*out = *in
	if in.ReconcilerPeriod != nil {
		in, out := &in.ReconcilerPeriod, &out.ReconcilerPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceControllerConfig.
func (in *NamespaceControllerConfig) DeepCopy() *NamespaceControllerConfig {
	if in == nil {
		return nil
	}
	out := new(NamespaceControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeControllerConfig) DeepCopyInto(out *NodeControllerConfig) {
	*out = *in
	if in.ReconcilerPeriod != nil {
		in, out := &in.ReconcilerPeriod, &out.ReconcilerPeriod
		*out = new(metav1.Duration)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeControllerConfig.
func (in *NodeControllerConfig) DeepCopy() *NodeControllerConfig {
	if in == nil {
		return nil
	}
	out := new(NodeControllerConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTopology) DeepCopyInto(out *NodeTopology) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTopology.
func (in *NodeTopology) DeepCopy() *NodeTopology {
	if in == nil {
		return nil
	}
	out := new(NodeTopology)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeTopology) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTopologyList) DeepCopyInto(out *NodeTopologyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NodeTopology, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTopologyList.
func (in *NodeTopologyList) DeepCopy() *NodeTopologyList {
	if in == nil {
		return nil
	}
	out := new(NodeTopologyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NodeTopologyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTopologySpec) DeepCopyInto(out *NodeTopologySpec) {
	*out = *in
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]*Device, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Device)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTopologySpec.
func (in *NodeTopologySpec) DeepCopy() *NodeTopologySpec {
	if in == nil {
		return nil
	}
	out := new(NodeTopologySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Vlan) DeepCopyInto(out *Vlan) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Vlan.
func (in *Vlan) DeepCopy() *Vlan {
	if in == nil {
		return nil
	}
	out := new(Vlan)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkLoad) DeepCopyInto(out *WorkLoad) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkLoad.
func (in *WorkLoad) DeepCopy() *WorkLoad {
	if in == nil {
		return nil
	}
	out := new(WorkLoad)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkLoad) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkLoadList) DeepCopyInto(out *WorkLoadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WorkLoad, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkLoadList.
func (in *WorkLoadList) DeepCopy() *WorkLoadList {
	if in == nil {
		return nil
	}
	out := new(WorkLoadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkLoadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkLoadSpec) DeepCopyInto(out *WorkLoadSpec) {
	*out = *in
	if in.Vlans != nil {
		in, out := &in.Vlans, &out.Vlans
		*out = make([]*Vlan, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Vlan)
				**out = **in
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkLoadSpec.
func (in *WorkLoadSpec) DeepCopy() *WorkLoadSpec {
	if in == nil {
		return nil
	}
	out := new(WorkLoadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkLoadStatus) DeepCopyInto(out *WorkLoadStatus) {
	*out = *in
	in.RunningConfig.DeepCopyInto(&out.RunningConfig)
	if in.Devices != nil {
		in, out := &in.Devices, &out.Devices
		*out = make([]*Device, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Device)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkLoadStatus.
func (in *WorkLoadStatus) DeepCopy() *WorkLoadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkLoadStatus)
	in.DeepCopyInto(out)
	return out
}

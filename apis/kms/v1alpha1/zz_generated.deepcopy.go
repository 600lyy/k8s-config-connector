//go:build !ignore_autogenerated

// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/GoogleCloudPlatform/k8s-config-connector/apis/refs/v1beta1"
	k8sv1alpha1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/apis/k8s/v1alpha1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AutokeyConfig) DeepCopyInto(out *AutokeyConfig) {
	*out = *in
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.KeyProject != nil {
		in, out := &in.KeyProject, &out.KeyProject
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AutokeyConfig.
func (in *AutokeyConfig) DeepCopy() *AutokeyConfig {
	if in == nil {
		return nil
	}
	out := new(AutokeyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSAutokeyConfig) DeepCopyInto(out *KMSAutokeyConfig) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSAutokeyConfig.
func (in *KMSAutokeyConfig) DeepCopy() *KMSAutokeyConfig {
	if in == nil {
		return nil
	}
	out := new(KMSAutokeyConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KMSAutokeyConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSAutokeyConfigList) DeepCopyInto(out *KMSAutokeyConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]KMSAutokeyConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSAutokeyConfigList.
func (in *KMSAutokeyConfigList) DeepCopy() *KMSAutokeyConfigList {
	if in == nil {
		return nil
	}
	out := new(KMSAutokeyConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KMSAutokeyConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSAutokeyConfigObservedState) DeepCopyInto(out *KMSAutokeyConfigObservedState) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSAutokeyConfigObservedState.
func (in *KMSAutokeyConfigObservedState) DeepCopy() *KMSAutokeyConfigObservedState {
	if in == nil {
		return nil
	}
	out := new(KMSAutokeyConfigObservedState)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSAutokeyConfigParent) DeepCopyInto(out *KMSAutokeyConfigParent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSAutokeyConfigParent.
func (in *KMSAutokeyConfigParent) DeepCopy() *KMSAutokeyConfigParent {
	if in == nil {
		return nil
	}
	out := new(KMSAutokeyConfigParent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSAutokeyConfigRef) DeepCopyInto(out *KMSAutokeyConfigRef) {
	*out = *in
	if in.parent != nil {
		in, out := &in.parent, &out.parent
		*out = new(KMSAutokeyConfigParent)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSAutokeyConfigRef.
func (in *KMSAutokeyConfigRef) DeepCopy() *KMSAutokeyConfigRef {
	if in == nil {
		return nil
	}
	out := new(KMSAutokeyConfigRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSAutokeyConfigSpec) DeepCopyInto(out *KMSAutokeyConfigSpec) {
	*out = *in
	if in.FolderRef != nil {
		in, out := &in.FolderRef, &out.FolderRef
		*out = new(v1beta1.FolderRef)
		**out = **in
	}
	if in.KeyProjectRef != nil {
		in, out := &in.KeyProjectRef, &out.KeyProjectRef
		*out = new(v1beta1.ProjectRef)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSAutokeyConfigSpec.
func (in *KMSAutokeyConfigSpec) DeepCopy() *KMSAutokeyConfigSpec {
	if in == nil {
		return nil
	}
	out := new(KMSAutokeyConfigSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KMSAutokeyConfigStatus) DeepCopyInto(out *KMSAutokeyConfigStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]k8sv1alpha1.Condition, len(*in))
		copy(*out, *in)
	}
	if in.ObservedGeneration != nil {
		in, out := &in.ObservedGeneration, &out.ObservedGeneration
		*out = new(int64)
		**out = **in
	}
	if in.ExternalRef != nil {
		in, out := &in.ExternalRef, &out.ExternalRef
		*out = new(string)
		**out = **in
	}
	if in.ObservedState != nil {
		in, out := &in.ObservedState, &out.ObservedState
		*out = new(KMSAutokeyConfigObservedState)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KMSAutokeyConfigStatus.
func (in *KMSAutokeyConfigStatus) DeepCopy() *KMSAutokeyConfigStatus {
	if in == nil {
		return nil
	}
	out := new(KMSAutokeyConfigStatus)
	in.DeepCopyInto(out)
	return out
}
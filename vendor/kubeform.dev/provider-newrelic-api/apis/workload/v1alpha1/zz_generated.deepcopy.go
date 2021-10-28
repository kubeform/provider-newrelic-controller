//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright AppsCode Inc. and Contributors

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

package v1alpha1

import (
	v1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Workload) DeepCopyInto(out *Workload) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Workload.
func (in *Workload) DeepCopy() *Workload {
	if in == nil {
		return nil
	}
	out := new(Workload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Workload) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadList) DeepCopyInto(out *WorkloadList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Workload, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadList.
func (in *WorkloadList) DeepCopy() *WorkloadList {
	if in == nil {
		return nil
	}
	out := new(WorkloadList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WorkloadList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSpec) DeepCopyInto(out *WorkloadSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(WorkloadSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	if in.BackendRef != nil {
		in, out := &in.BackendRef, &out.BackendRef
		*out = new(v1.LocalObjectReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSpec.
func (in *WorkloadSpec) DeepCopy() *WorkloadSpec {
	if in == nil {
		return nil
	}
	out := new(WorkloadSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSpecEntitySearchQuery) DeepCopyInto(out *WorkloadSpecEntitySearchQuery) {
	*out = *in
	if in.Query != nil {
		in, out := &in.Query, &out.Query
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSpecEntitySearchQuery.
func (in *WorkloadSpecEntitySearchQuery) DeepCopy() *WorkloadSpecEntitySearchQuery {
	if in == nil {
		return nil
	}
	out := new(WorkloadSpecEntitySearchQuery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadSpecResource) DeepCopyInto(out *WorkloadSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(int64)
		**out = **in
	}
	if in.CompositeEntitySearchQuery != nil {
		in, out := &in.CompositeEntitySearchQuery, &out.CompositeEntitySearchQuery
		*out = new(string)
		**out = **in
	}
	if in.EntityGuids != nil {
		in, out := &in.EntityGuids, &out.EntityGuids
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.EntitySearchQuery != nil {
		in, out := &in.EntitySearchQuery, &out.EntitySearchQuery
		*out = make([]WorkloadSpecEntitySearchQuery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Guid != nil {
		in, out := &in.Guid, &out.Guid
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Permalink != nil {
		in, out := &in.Permalink, &out.Permalink
		*out = new(string)
		**out = **in
	}
	if in.ScopeAccountIDS != nil {
		in, out := &in.ScopeAccountIDS, &out.ScopeAccountIDS
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
	if in.WorkloadID != nil {
		in, out := &in.WorkloadID, &out.WorkloadID
		*out = new(int64)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadSpecResource.
func (in *WorkloadSpecResource) DeepCopy() *WorkloadSpecResource {
	if in == nil {
		return nil
	}
	out := new(WorkloadSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkloadStatus) DeepCopyInto(out *WorkloadStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]apiv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkloadStatus.
func (in *WorkloadStatus) DeepCopy() *WorkloadStatus {
	if in == nil {
		return nil
	}
	out := new(WorkloadStatus)
	in.DeepCopyInto(out)
	return out
}

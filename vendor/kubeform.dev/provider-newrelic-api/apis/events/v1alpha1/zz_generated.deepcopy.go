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
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1 "kmodules.xyz/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToMetricsRule) DeepCopyInto(out *ToMetricsRule) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToMetricsRule.
func (in *ToMetricsRule) DeepCopy() *ToMetricsRule {
	if in == nil {
		return nil
	}
	out := new(ToMetricsRule)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ToMetricsRule) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToMetricsRuleList) DeepCopyInto(out *ToMetricsRuleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ToMetricsRule, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToMetricsRuleList.
func (in *ToMetricsRuleList) DeepCopy() *ToMetricsRuleList {
	if in == nil {
		return nil
	}
	out := new(ToMetricsRuleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ToMetricsRuleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToMetricsRuleSpec) DeepCopyInto(out *ToMetricsRuleSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(ToMetricsRuleSpecResource)
		(*in).DeepCopyInto(*out)
	}
	in.Resource.DeepCopyInto(&out.Resource)
	out.ProviderRef = in.ProviderRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToMetricsRuleSpec.
func (in *ToMetricsRuleSpec) DeepCopy() *ToMetricsRuleSpec {
	if in == nil {
		return nil
	}
	out := new(ToMetricsRuleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToMetricsRuleSpecResource) DeepCopyInto(out *ToMetricsRuleSpecResource) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(int64)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Nrql != nil {
		in, out := &in.Nrql, &out.Nrql
		*out = new(string)
		**out = **in
	}
	if in.RuleID != nil {
		in, out := &in.RuleID, &out.RuleID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToMetricsRuleSpecResource.
func (in *ToMetricsRuleSpecResource) DeepCopy() *ToMetricsRuleSpecResource {
	if in == nil {
		return nil
	}
	out := new(ToMetricsRuleSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ToMetricsRuleStatus) DeepCopyInto(out *ToMetricsRuleStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ToMetricsRuleStatus.
func (in *ToMetricsRuleStatus) DeepCopy() *ToMetricsRuleStatus {
	if in == nil {
		return nil
	}
	out := new(ToMetricsRuleStatus)
	in.DeepCopyInto(out)
	return out
}

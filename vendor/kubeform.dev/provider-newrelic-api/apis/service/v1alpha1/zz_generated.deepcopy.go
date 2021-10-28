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
func (in *Level) DeepCopyInto(out *Level) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Level.
func (in *Level) DeepCopy() *Level {
	if in == nil {
		return nil
	}
	out := new(Level)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Level) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelList) DeepCopyInto(out *LevelList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Level, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelList.
func (in *LevelList) DeepCopy() *LevelList {
	if in == nil {
		return nil
	}
	out := new(LevelList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LevelList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpec) DeepCopyInto(out *LevelSpec) {
	*out = *in
	if in.State != nil {
		in, out := &in.State, &out.State
		*out = new(LevelSpecResource)
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpec.
func (in *LevelSpec) DeepCopy() *LevelSpec {
	if in == nil {
		return nil
	}
	out := new(LevelSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecEvents) DeepCopyInto(out *LevelSpecEvents) {
	*out = *in
	if in.AccountID != nil {
		in, out := &in.AccountID, &out.AccountID
		*out = new(int64)
		**out = **in
	}
	if in.BadEvents != nil {
		in, out := &in.BadEvents, &out.BadEvents
		*out = new(LevelSpecEventsBadEvents)
		(*in).DeepCopyInto(*out)
	}
	if in.GoodEvents != nil {
		in, out := &in.GoodEvents, &out.GoodEvents
		*out = new(LevelSpecEventsGoodEvents)
		(*in).DeepCopyInto(*out)
	}
	if in.ValidEvents != nil {
		in, out := &in.ValidEvents, &out.ValidEvents
		*out = new(LevelSpecEventsValidEvents)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecEvents.
func (in *LevelSpecEvents) DeepCopy() *LevelSpecEvents {
	if in == nil {
		return nil
	}
	out := new(LevelSpecEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecEventsBadEvents) DeepCopyInto(out *LevelSpecEventsBadEvents) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.Where != nil {
		in, out := &in.Where, &out.Where
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecEventsBadEvents.
func (in *LevelSpecEventsBadEvents) DeepCopy() *LevelSpecEventsBadEvents {
	if in == nil {
		return nil
	}
	out := new(LevelSpecEventsBadEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecEventsGoodEvents) DeepCopyInto(out *LevelSpecEventsGoodEvents) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.Where != nil {
		in, out := &in.Where, &out.Where
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecEventsGoodEvents.
func (in *LevelSpecEventsGoodEvents) DeepCopy() *LevelSpecEventsGoodEvents {
	if in == nil {
		return nil
	}
	out := new(LevelSpecEventsGoodEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecEventsValidEvents) DeepCopyInto(out *LevelSpecEventsValidEvents) {
	*out = *in
	if in.From != nil {
		in, out := &in.From, &out.From
		*out = new(string)
		**out = **in
	}
	if in.Where != nil {
		in, out := &in.Where, &out.Where
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecEventsValidEvents.
func (in *LevelSpecEventsValidEvents) DeepCopy() *LevelSpecEventsValidEvents {
	if in == nil {
		return nil
	}
	out := new(LevelSpecEventsValidEvents)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecObjective) DeepCopyInto(out *LevelSpecObjective) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Name != nil {
		in, out := &in.Name, &out.Name
		*out = new(string)
		**out = **in
	}
	if in.Target != nil {
		in, out := &in.Target, &out.Target
		*out = new(float64)
		**out = **in
	}
	if in.TimeWindow != nil {
		in, out := &in.TimeWindow, &out.TimeWindow
		*out = new(LevelSpecObjectiveTimeWindow)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecObjective.
func (in *LevelSpecObjective) DeepCopy() *LevelSpecObjective {
	if in == nil {
		return nil
	}
	out := new(LevelSpecObjective)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecObjectiveTimeWindow) DeepCopyInto(out *LevelSpecObjectiveTimeWindow) {
	*out = *in
	if in.Rolling != nil {
		in, out := &in.Rolling, &out.Rolling
		*out = new(LevelSpecObjectiveTimeWindowRolling)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecObjectiveTimeWindow.
func (in *LevelSpecObjectiveTimeWindow) DeepCopy() *LevelSpecObjectiveTimeWindow {
	if in == nil {
		return nil
	}
	out := new(LevelSpecObjectiveTimeWindow)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecObjectiveTimeWindowRolling) DeepCopyInto(out *LevelSpecObjectiveTimeWindowRolling) {
	*out = *in
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(int64)
		**out = **in
	}
	if in.Unit != nil {
		in, out := &in.Unit, &out.Unit
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecObjectiveTimeWindowRolling.
func (in *LevelSpecObjectiveTimeWindowRolling) DeepCopy() *LevelSpecObjectiveTimeWindowRolling {
	if in == nil {
		return nil
	}
	out := new(LevelSpecObjectiveTimeWindowRolling)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelSpecResource) DeepCopyInto(out *LevelSpecResource) {
	*out = *in
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = new(LevelSpecEvents)
		(*in).DeepCopyInto(*out)
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
	if in.Objective != nil {
		in, out := &in.Objective, &out.Objective
		*out = make([]LevelSpecObjective, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.SliID != nil {
		in, out := &in.SliID, &out.SliID
		*out = new(string)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelSpecResource.
func (in *LevelSpecResource) DeepCopy() *LevelSpecResource {
	if in == nil {
		return nil
	}
	out := new(LevelSpecResource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LevelStatus) DeepCopyInto(out *LevelStatus) {
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

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LevelStatus.
func (in *LevelStatus) DeepCopy() *LevelStatus {
	if in == nil {
		return nil
	}
	out := new(LevelStatus)
	in.DeepCopyInto(out)
	return out
}

// +build !ignore_autogenerated

/*
Copyright 2017 The Kubernetes Authors.

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

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
//
// Deprecated: deepcopy registration will go away when static deepcopy is fully implemented.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*Instance).DeepCopyInto(out.(*Instance))
			return nil
		}, InType: reflect.TypeOf(&Instance{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InstanceList).DeepCopyInto(out.(*InstanceList))
			return nil
		}, InType: reflect.TypeOf(&InstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InstanceSpec).DeepCopyInto(out.(*InstanceSpec))
			return nil
		}, InType: reflect.TypeOf(&InstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InstanceStatus).DeepCopyInto(out.(*InstanceStatus))
			return nil
		}, InType: reflect.TypeOf(&InstanceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: func(in interface{}, out interface{}, c *conversion.Cloner) error {
			in.(*InstanceStatusManifest).DeepCopyInto(out.(*InstanceStatusManifest))
			return nil
		}, InType: reflect.TypeOf(&InstanceStatusManifest{})},
	)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Instance) DeepCopyInto(out *Instance) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		if *in == nil {
			*out = nil
		} else {
			*out = new(InstanceSpec)
			**out = **in
		}
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		if *in == nil {
			*out = nil
		} else {
			*out = new(InstanceStatus)
			(*in).DeepCopyInto(*out)
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Instance.
func (in *Instance) DeepCopy() *Instance {
	if in == nil {
		return nil
	}
	out := new(Instance)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Instance) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceList) DeepCopyInto(out *InstanceList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Instance, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceList.
func (in *InstanceList) DeepCopy() *InstanceList {
	if in == nil {
		return nil
	}
	out := new(InstanceList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *InstanceList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceSpec) DeepCopyInto(out *InstanceSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceSpec.
func (in *InstanceSpec) DeepCopy() *InstanceSpec {
	if in == nil {
		return nil
	}
	out := new(InstanceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatus) DeepCopyInto(out *InstanceStatus) {
	*out = *in
	if in.Converge != nil {
		in, out := &in.Converge, &out.Converge
		if *in == nil {
			*out = nil
		} else {
			*out = new(InstanceStatusManifest)
			**out = **in
		}
	}
	if in.DryRun != nil {
		in, out := &in.DryRun, &out.DryRun
		if *in == nil {
			*out = nil
		} else {
			*out = new(InstanceStatusManifest)
			**out = **in
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatus.
func (in *InstanceStatus) DeepCopy() *InstanceStatus {
	if in == nil {
		return nil
	}
	out := new(InstanceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstanceStatusManifest) DeepCopyInto(out *InstanceStatusManifest) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstanceStatusManifest.
func (in *InstanceStatusManifest) DeepCopy() *InstanceStatusManifest {
	if in == nil {
		return nil
	}
	out := new(InstanceStatusManifest)
	in.DeepCopyInto(out)
	return out
}

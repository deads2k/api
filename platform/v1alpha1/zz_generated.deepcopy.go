//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by codegen. DO NOT EDIT.

package v1alpha1

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveBundleDeployment) DeepCopyInto(out *ActiveBundleDeployment) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveBundleDeployment.
func (in *ActiveBundleDeployment) DeepCopy() *ActiveBundleDeployment {
	if in == nil {
		return nil
	}
	out := new(ActiveBundleDeployment)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Package) DeepCopyInto(out *Package) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Package.
func (in *Package) DeepCopy() *Package {
	if in == nil {
		return nil
	}
	out := new(Package)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperator) DeepCopyInto(out *PlatformOperator) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperator.
func (in *PlatformOperator) DeepCopy() *PlatformOperator {
	if in == nil {
		return nil
	}
	out := new(PlatformOperator)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformOperator) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperatorList) DeepCopyInto(out *PlatformOperatorList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PlatformOperator, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperatorList.
func (in *PlatformOperatorList) DeepCopy() *PlatformOperatorList {
	if in == nil {
		return nil
	}
	out := new(PlatformOperatorList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PlatformOperatorList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperatorSpec) DeepCopyInto(out *PlatformOperatorSpec) {
	*out = *in
	out.Package = in.Package
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperatorSpec.
func (in *PlatformOperatorSpec) DeepCopy() *PlatformOperatorSpec {
	if in == nil {
		return nil
	}
	out := new(PlatformOperatorSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlatformOperatorStatus) DeepCopyInto(out *PlatformOperatorStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	out.ActiveBundleDeployment = in.ActiveBundleDeployment
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlatformOperatorStatus.
func (in *PlatformOperatorStatus) DeepCopy() *PlatformOperatorStatus {
	if in == nil {
		return nil
	}
	out := new(PlatformOperatorStatus)
	in.DeepCopyInto(out)
	return out
}

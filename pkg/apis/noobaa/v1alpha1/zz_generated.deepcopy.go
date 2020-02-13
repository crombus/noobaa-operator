// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AWSS3Spec) DeepCopyInto(out *AWSS3Spec) {
	*out = *in
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AWSS3Spec.
func (in *AWSS3Spec) DeepCopy() *AWSS3Spec {
	if in == nil {
		return nil
	}
	out := new(AWSS3Spec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AccountsStatus) DeepCopyInto(out *AccountsStatus) {
	*out = *in
	out.Admin = in.Admin
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AccountsStatus.
func (in *AccountsStatus) DeepCopy() *AccountsStatus {
	if in == nil {
		return nil
	}
	out := new(AccountsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AzureBlobSpec) DeepCopyInto(out *AzureBlobSpec) {
	*out = *in
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AzureBlobSpec.
func (in *AzureBlobSpec) DeepCopy() *AzureBlobSpec {
	if in == nil {
		return nil
	}
	out := new(AzureBlobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStore) DeepCopyInto(out *BackingStore) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStore.
func (in *BackingStore) DeepCopy() *BackingStore {
	if in == nil {
		return nil
	}
	out := new(BackingStore)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackingStore) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStoreList) DeepCopyInto(out *BackingStoreList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BackingStore, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStoreList.
func (in *BackingStoreList) DeepCopy() *BackingStoreList {
	if in == nil {
		return nil
	}
	out := new(BackingStoreList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BackingStoreList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStoreMode) DeepCopyInto(out *BackingStoreMode) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStoreMode.
func (in *BackingStoreMode) DeepCopy() *BackingStoreMode {
	if in == nil {
		return nil
	}
	out := new(BackingStoreMode)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStorePhaseInfo) DeepCopyInto(out *BackingStorePhaseInfo) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStorePhaseInfo.
func (in *BackingStorePhaseInfo) DeepCopy() *BackingStorePhaseInfo {
	if in == nil {
		return nil
	}
	out := new(BackingStorePhaseInfo)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStoreSpec) DeepCopyInto(out *BackingStoreSpec) {
	*out = *in
	if in.AWSS3 != nil {
		in, out := &in.AWSS3, &out.AWSS3
		*out = new(AWSS3Spec)
		**out = **in
	}
	if in.S3Compatible != nil {
		in, out := &in.S3Compatible, &out.S3Compatible
		*out = new(S3CompatibleSpec)
		**out = **in
	}
	if in.AzureBlob != nil {
		in, out := &in.AzureBlob, &out.AzureBlob
		*out = new(AzureBlobSpec)
		**out = **in
	}
	if in.GoogleCloudStorage != nil {
		in, out := &in.GoogleCloudStorage, &out.GoogleCloudStorage
		*out = new(GoogleCloudStorageSpec)
		**out = **in
	}
	if in.PVPool != nil {
		in, out := &in.PVPool, &out.PVPool
		*out = new(PVPoolSpec)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStoreSpec.
func (in *BackingStoreSpec) DeepCopy() *BackingStoreSpec {
	if in == nil {
		return nil
	}
	out := new(BackingStoreSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BackingStoreStatus) DeepCopyInto(out *BackingStoreStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	out.Mode = in.Mode
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BackingStoreStatus.
func (in *BackingStoreStatus) DeepCopy() *BackingStoreStatus {
	if in == nil {
		return nil
	}
	out := new(BackingStoreStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClass) DeepCopyInto(out *BucketClass) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClass.
func (in *BucketClass) DeepCopy() *BucketClass {
	if in == nil {
		return nil
	}
	out := new(BucketClass)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketClass) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassList) DeepCopyInto(out *BucketClassList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]BucketClass, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassList.
func (in *BucketClassList) DeepCopy() *BucketClassList {
	if in == nil {
		return nil
	}
	out := new(BucketClassList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *BucketClassList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassSpec) DeepCopyInto(out *BucketClassSpec) {
	*out = *in
	in.PlacementPolicy.DeepCopyInto(&out.PlacementPolicy)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassSpec.
func (in *BucketClassSpec) DeepCopy() *BucketClassSpec {
	if in == nil {
		return nil
	}
	out := new(BucketClassSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BucketClassStatus) DeepCopyInto(out *BucketClassStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new BucketClassStatus.
func (in *BucketClassStatus) DeepCopy() *BucketClassStatus {
	if in == nil {
		return nil
	}
	out := new(BucketClassStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointsSpec) DeepCopyInto(out *EndpointsSpec) {
	*out = *in
	if in.AdditionalVirtualHosts != nil {
		in, out := &in.AdditionalVirtualHosts, &out.AdditionalVirtualHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointsSpec.
func (in *EndpointsSpec) DeepCopy() *EndpointsSpec {
	if in == nil {
		return nil
	}
	out := new(EndpointsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EndpointsStatus) DeepCopyInto(out *EndpointsStatus) {
	*out = *in
	if in.VirtualHosts != nil {
		in, out := &in.VirtualHosts, &out.VirtualHosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EndpointsStatus.
func (in *EndpointsStatus) DeepCopy() *EndpointsStatus {
	if in == nil {
		return nil
	}
	out := new(EndpointsStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GoogleCloudStorageSpec) DeepCopyInto(out *GoogleCloudStorageSpec) {
	*out = *in
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GoogleCloudStorageSpec.
func (in *GoogleCloudStorageSpec) DeepCopy() *GoogleCloudStorageSpec {
	if in == nil {
		return nil
	}
	out := new(GoogleCloudStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaa) DeepCopyInto(out *NooBaa) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaa.
func (in *NooBaa) DeepCopy() *NooBaa {
	if in == nil {
		return nil
	}
	out := new(NooBaa)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NooBaa) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaList) DeepCopyInto(out *NooBaaList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]NooBaa, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaList.
func (in *NooBaaList) DeepCopy() *NooBaaList {
	if in == nil {
		return nil
	}
	out := new(NooBaaList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *NooBaaList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaSpec) DeepCopyInto(out *NooBaaSpec) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(string)
		**out = **in
	}
	if in.DBImage != nil {
		in, out := &in.DBImage, &out.DBImage
		*out = new(string)
		**out = **in
	}
	if in.CoreResources != nil {
		in, out := &in.CoreResources, &out.CoreResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.DBResources != nil {
		in, out := &in.DBResources, &out.DBResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.DBVolumeResources != nil {
		in, out := &in.DBVolumeResources, &out.DBVolumeResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	if in.DBStorageClass != nil {
		in, out := &in.DBStorageClass, &out.DBStorageClass
		*out = new(string)
		**out = **in
	}
	if in.PVPoolDefaultStorageClass != nil {
		in, out := &in.PVPoolDefaultStorageClass, &out.PVPoolDefaultStorageClass
		*out = new(string)
		**out = **in
	}
	if in.Tolerations != nil {
		in, out := &in.Tolerations, &out.Tolerations
		*out = make([]corev1.Toleration, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(corev1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	if in.ImagePullSecret != nil {
		in, out := &in.ImagePullSecret, &out.ImagePullSecret
		*out = new(corev1.LocalObjectReference)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = new(EndpointsSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.JoinSecret != nil {
		in, out := &in.JoinSecret, &out.JoinSecret
		*out = new(corev1.SecretReference)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaSpec.
func (in *NooBaaSpec) DeepCopy() *NooBaaSpec {
	if in == nil {
		return nil
	}
	out := new(NooBaaSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NooBaaStatus) DeepCopyInto(out *NooBaaStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.Accounts != nil {
		in, out := &in.Accounts, &out.Accounts
		*out = new(AccountsStatus)
		**out = **in
	}
	if in.Services != nil {
		in, out := &in.Services, &out.Services
		*out = new(ServicesStatus)
		(*in).DeepCopyInto(*out)
	}
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = new(EndpointsStatus)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NooBaaStatus.
func (in *NooBaaStatus) DeepCopy() *NooBaaStatus {
	if in == nil {
		return nil
	}
	out := new(NooBaaStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PVPoolSpec) DeepCopyInto(out *PVPoolSpec) {
	*out = *in
	if in.VolumeResources != nil {
		in, out := &in.VolumeResources, &out.VolumeResources
		*out = new(corev1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PVPoolSpec.
func (in *PVPoolSpec) DeepCopy() *PVPoolSpec {
	if in == nil {
		return nil
	}
	out := new(PVPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PlacementPolicy) DeepCopyInto(out *PlacementPolicy) {
	*out = *in
	if in.Tiers != nil {
		in, out := &in.Tiers, &out.Tiers
		*out = make([]Tier, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PlacementPolicy.
func (in *PlacementPolicy) DeepCopy() *PlacementPolicy {
	if in == nil {
		return nil
	}
	out := new(PlacementPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3CompatibleSpec) DeepCopyInto(out *S3CompatibleSpec) {
	*out = *in
	out.Secret = in.Secret
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3CompatibleSpec.
func (in *S3CompatibleSpec) DeepCopy() *S3CompatibleSpec {
	if in == nil {
		return nil
	}
	out := new(S3CompatibleSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServiceStatus) DeepCopyInto(out *ServiceStatus) {
	*out = *in
	if in.NodePorts != nil {
		in, out := &in.NodePorts, &out.NodePorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.PodPorts != nil {
		in, out := &in.PodPorts, &out.PodPorts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InternalIP != nil {
		in, out := &in.InternalIP, &out.InternalIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.InternalDNS != nil {
		in, out := &in.InternalDNS, &out.InternalDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.ExternalDNS != nil {
		in, out := &in.ExternalDNS, &out.ExternalDNS
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServiceStatus.
func (in *ServiceStatus) DeepCopy() *ServiceStatus {
	if in == nil {
		return nil
	}
	out := new(ServiceStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ServicesStatus) DeepCopyInto(out *ServicesStatus) {
	*out = *in
	in.ServiceMgmt.DeepCopyInto(&out.ServiceMgmt)
	in.ServiceS3.DeepCopyInto(&out.ServiceS3)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ServicesStatus.
func (in *ServicesStatus) DeepCopy() *ServicesStatus {
	if in == nil {
		return nil
	}
	out := new(ServicesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tier) DeepCopyInto(out *Tier) {
	*out = *in
	if in.BackingStores != nil {
		in, out := &in.BackingStores, &out.BackingStores
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tier.
func (in *Tier) DeepCopy() *Tier {
	if in == nil {
		return nil
	}
	out := new(Tier)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *UserStatus) DeepCopyInto(out *UserStatus) {
	*out = *in
	out.SecretRef = in.SecretRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new UserStatus.
func (in *UserStatus) DeepCopy() *UserStatus {
	if in == nil {
		return nil
	}
	out := new(UserStatus)
	in.DeepCopyInto(out)
	return out
}

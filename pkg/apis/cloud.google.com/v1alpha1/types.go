package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedCertificateList is a list of ManagedCertificate objects.
type ManagedCertificateList struct {
	metav1.TypeMeta `json:",inline"`
	// metdata is the standard list metadata.
	// +optional
	metav1.ListMeta `json:"metadata" protobuf:"bytes,1,opt,name=metadata"`

	// items is the list of managed certificate objects.
	Items []ManagedCertificate `json:"items" protobuf:"bytes,2,rep,name=items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ManagedCertificate configures for what domains the client requests managed certificate and with which Ingress they should be associated. It also provides the current status of the certficate.
type ManagedCertificate struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Specification of the managed certificate.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status.
	// +optional
	Spec ManagedCertificateSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`

	// Current information about the managed certificate.
	// +optional
	Status ManagedCertificateStatus `json:"status,omitempty" protobuf:"bytes,3,opt,name=status"`
}

// ManagedCertificateSpec configures for what domains the client requests managed certificate and with which Ingress it should be associated.
type ManagedCertificateSpec struct {
	// Specifies a list of domains populated by the user for which he requests a managed certificate.
	Domains []string `json:"domains" protobuf:"bytes,2,rep,name=domains"`

	// Output only: specifies a final list of domains for which a certificate is created. This final list is build by merging the contents of the Domains field with domains extracted from matching Ingress configurations.
	ExpandedDomains []string `json:"expandedDomains" protobuf:"bytes,2,rep,name=expandedDomains"`
}

// ManagedCertificateStatus provides the current state of the certificate.
type ManagedCertificateStatus struct {
	// Specifies the status of the managed certificate.
	// +optional
	CertificateStatus string `json:"certificateStatus,omitempty" protobuf:"bytes,1,opt,name=certificateStatus"`

	// Specifies the statuses of certificate provisioning for domains selected by the user.
	DomainStatus []DomainStatus `json:"domainStatus" protobuf:"bytes,2,rep,name=domainStatus"`

	// Specifies the name of the provisioned managed certificate.
	// +optional
	CertificateName string `json:"certificateName,omitempty" protobuf:"bytes,3,opt,name=certificateName"`
}

// DomainStatus is a pair which associates domain name with status of certificate provisioning for this domain.
type DomainStatus struct {
	// The domain name.
	Domain string `json:"domain" protobuf:"bytes,1,name=domain"`

	// The status.
	Status string `json:"status" protobuf:"bytes,2,name=status"`
}
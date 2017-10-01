package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	clusterv1alpha1 "github.com/jetstack/tarmak/pkg/apis/cluster/v1alpha1"
)

// +genclient=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=configs

type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	CurrentCluster string `json:"currentCluster,omitempty"` // <environmentName>-<clusterName>

	Contact string `json:"contact,omitempty"`
	Project string `json:"project,omitempty"`

	Clusters     []clusterv1alpha1.Cluster `json:"clusters,omitempty"`
	Providers    []Provider                `json:"providers,omitempty"`
	Environments []Environment             `json:"environments,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ConfigList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Config `json:"items"`
}

// +genclient=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=providers

type Provider struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Amazon *ProviderAmazon `json:"amazon,omitempty"`
	GCP    *ProviderGCP    `json:"gcp,omitempty"`
	Azure  *ProviderAzure  `json:"azure,omitempty"`
}

type ProviderAmazon struct {
	VaultPath         string   `json:"vaultPath,omitempty"`
	AllowedAccountIDs []string `json:"allowedAccountIDs,omitempty"`
	Profile           string   `json:"profile,omitempty"`
	BucketPrefix      string   `json:"bucketPrefix,omitempty"`
	KeyName           string   `json:"keyName,omitempty"`

	PublicZone         string `json:"publicZone,omitempty"`
	PublicHostedZoneID string `json:"publicHostedZoneID, omitempty"`
}

type ProviderGCP struct {
	Project string `json:"project,omitempty"`
}

type ProviderAzure struct {
	SubscriptionID string `json:"subscriptionID,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type ProviderList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Provider `json:"items"`
}

// +genclient=true
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +resource:path=environments

type Environment struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Provider string `json:"provider,omitempty"`

	Contact string `json:"contact,omitempty"`
	Project string `json:"project,omitempty"`

	Location    string               `json:"location,omitempty"`
	SSH         *clusterv1alpha1.SSH `json:"ssh,omitempty"`
	PrivateZone string               `json:"privateZone,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type EnvironmentList struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Items []Environment `json:"items"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type Image struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	BaseImage string `json:"baseImage,omitempty"`
	Location  string `json:"location,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

type InstanceState struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	InstanceID   string `json:"instanceID,omitempty"`
	InstancePool string `json:"instanceID,omitempty"`

	Spec   *InstanceStateSpec   `json:"spec,omitempty"`
	Status *InstanceStateStatus `json:"status,omitempty"`
}

type InstanceStateSpec struct {
	ConvergeHash string `json:"convergeHash,omitempty"`
	DryRunPath   string `json:"dryRunPath,omitempty"`
	DryRunHash   string `json:"dryRunHash,omitempty"`
}

type InstanceStateStatus struct {
	Converge *InstanceStateStatusManifest `json:"converge,omitempty"`
	DryRun   *InstanceStateStatusManifest `json:"dryRun,omitempty"`
}

type InstanceStateStatusManifest struct {
	State string `json:"state,omitempty"`
	Hash  string `json:"hash,omitempty"`
}

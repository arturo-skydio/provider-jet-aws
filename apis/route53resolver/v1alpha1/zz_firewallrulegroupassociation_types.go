/*
Copyright 2021 The Crossplane Authors.

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

// Code generated by terrajet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type FirewallRuleGroupAssociationObservation struct {
	Arn *string `json:"arn,omitempty" tf:"arn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	TagsAll map[string]*string `json:"tagsAll,omitempty" tf:"tags_all,omitempty"`
}

type FirewallRuleGroupAssociationParameters struct {

	// +kubebuilder:validation:Required
	FirewallRuleGroupID *string `json:"firewallRuleGroupId" tf:"firewall_rule_group_id,omitempty"`

	// +kubebuilder:validation:Optional
	MutationProtection *string `json:"mutationProtection,omitempty" tf:"mutation_protection,omitempty"`

	// +kubebuilder:validation:Required
	Name *string `json:"name" tf:"name,omitempty"`

	// +kubebuilder:validation:Required
	Priority *int64 `json:"priority" tf:"priority,omitempty"`

	// Region is the region you'd like your resource to be created in.
	// +terrajet:crd:field:TFTag=-
	// +kubebuilder:validation:Required
	Region *string `json:"region" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags map[string]*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +crossplane:generate:reference:type=github.com/crossplane-contrib/provider-jet-aws/apis/ec2/v1alpha2.VPC
	// +crossplane:generate:reference:refFieldName=VpcIdRef
	// +crossplane:generate:reference:selectorFieldName=VpcIdSelector
	// +kubebuilder:validation:Optional
	VPCID *string `json:"vpcId,omitempty" tf:"vpc_id,omitempty"`

	// +kubebuilder:validation:Optional
	VpcIdRef *v1.Reference `json:"vpcIdRef,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	VpcIdSelector *v1.Selector `json:"vpcIdSelector,omitempty" tf:"-"`
}

// FirewallRuleGroupAssociationSpec defines the desired state of FirewallRuleGroupAssociation
type FirewallRuleGroupAssociationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     FirewallRuleGroupAssociationParameters `json:"forProvider"`
}

// FirewallRuleGroupAssociationStatus defines the observed state of FirewallRuleGroupAssociation.
type FirewallRuleGroupAssociationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        FirewallRuleGroupAssociationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleGroupAssociation is the Schema for the FirewallRuleGroupAssociations API
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,awsjet}
type FirewallRuleGroupAssociation struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              FirewallRuleGroupAssociationSpec   `json:"spec"`
	Status            FirewallRuleGroupAssociationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// FirewallRuleGroupAssociationList contains a list of FirewallRuleGroupAssociations
type FirewallRuleGroupAssociationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []FirewallRuleGroupAssociation `json:"items"`
}

// Repository type metadata.
var (
	FirewallRuleGroupAssociation_Kind             = "FirewallRuleGroupAssociation"
	FirewallRuleGroupAssociation_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: FirewallRuleGroupAssociation_Kind}.String()
	FirewallRuleGroupAssociation_KindAPIVersion   = FirewallRuleGroupAssociation_Kind + "." + CRDGroupVersion.String()
	FirewallRuleGroupAssociation_GroupVersionKind = CRDGroupVersion.WithKind(FirewallRuleGroupAssociation_Kind)
)

func init() {
	SchemeBuilder.Register(&FirewallRuleGroupAssociation{}, &FirewallRuleGroupAssociationList{})
}

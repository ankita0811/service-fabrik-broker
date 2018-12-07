/*
Copyright 2018 The Service Fabrik Authors.

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

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// Context contains the data additional data regarding service/plan
type Context struct {
	Operator      *runtime.RawExtension `json:"operator,omitempty"`
	ServiceFabrik *runtime.RawExtension `json:"serviceFabrik,omitempty"`
}

// DashboardClient contains the data necessary to activate the Dashboard SSO feature for this service
type DashboardClient struct {
	ID          string `json:"id,omitempty"`
	Secret      string `json:"secret,omitempty"`
	RedirectURI string `json:"redirectUri,omitempty"`
}

// ServiceMetadata contains the metadata for the service
type ServiceMetadata struct {
	DisplayName         string `json:"displayName,omitempty"`
	LongDescription     string `json:"longDescription,omitempty"`
	ProviderDisplayName string `json:"providerDisplayName,omitempty"`
	DocumentationURL    string `json:"documentationUrl,omitempty"`
	SupportURL          string `json:"supportUrl,omitempty"`
	ImageURL            string `json:"imageUrl,omitempty"`
}

// SfServiceSpec defines the desired state of SfService
type SfServiceSpec struct {
	Name                string                `json:"name"`
	ID                  string                `json:"id"`
	Description         string                `json:"description"`
	Tags                []string              `json:"tags,omitempty"`
	Requires            []string              `json:"requires,omitempty"`
	Bindable            bool                  `json:"bindable"`
	InstanceRetrievable bool                  `json:"instanceRetrievable,omitempty"`
	BindingRetrievable  bool                  `json:"bindingRetrievable,omitempty"`
	Metadata            *runtime.RawExtension `json:"metadata,omitempty"`
	DashboardClient     DashboardClient       `json:"dashboardClient,omitempty"`
	PlanUpdatable       bool                  `json:"planUpdatable,omitempty"`
	RawContext          *runtime.RawExtension `json:"context,omitempty"`
}

// SfServiceStatus defines the observed state of SfService
type SfServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SfService is the Schema for the sfservices API
// +k8s:openapi-gen=true
type SfService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SfServiceSpec   `json:"spec,omitempty"`
	Status SfServiceStatus `json:"status,omitempty"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// SfServiceList contains a list of SfService
type SfServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SfService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SfService{}, &SfServiceList{})
}
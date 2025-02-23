package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DigitalCluster struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              CustomSpec `json:"spec"`
}

type CustomSpec struct {
	Name      string     `json:"name"`
	Region    string     `json:"region"`
	Version   string     `json:"version"`
	NodePools []NodePool `json:"nodePools"`
}

type NodePool struct {
	Size  string `json:"size"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
type DigitalClusterList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DigitalCluster `json:"items"`
}

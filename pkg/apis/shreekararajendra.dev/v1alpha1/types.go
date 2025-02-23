package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type DigitalCluster struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	spec CustomSpec
}

type CustomSpec struct {
	name       string
	region     string
	version    string
	node_pools []node_pool
}

type node_pool struct {
	size  string
	name  string
	count int
}

type DigitalClusterList struct {
	metav1.TypeMeta
	metav1.ObjectMeta
	Items []DigitalCluster
}

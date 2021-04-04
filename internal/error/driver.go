package error

import "fmt"

type ClusterExists struct{
	Name string
}

func (e *ClusterExists) Error() string {
	return fmt.Sprintf("cluster '%s' already exists", e.Name)
}

type KubernetesVersionType struct{}

func (e *KubernetesVersionType) Error() string {
	return "kubernetes-version has incorrect type, should be string"
}

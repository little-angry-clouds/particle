package error

type ClusterExists struct{}

func (e *ClusterExists) Error() string {
	return "cluster already exists"
}

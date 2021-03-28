package error

type InexistentCluster struct{}

func (e *InexistentCluster) Error() string {
	return "there's no cluster, so the cleanup will fail"
}

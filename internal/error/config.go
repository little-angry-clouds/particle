package error

type ParticleAlreadyInitialized struct{}

func (e *ParticleAlreadyInitialized) Error() string {
	return "particle already initialized"
}

type ParticleNotInitialized struct{}

func (e *ParticleNotInitialized) Error() string {
	return "particle is not initialized"
}

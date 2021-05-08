package config

type particleAlreadyInitialized struct{}

func (e *particleAlreadyInitialized) Error() string {
	return "particle already initialized"
}

type particleNotInitialized struct{}

func (e *particleNotInitialized) Error() string {
	return "particle is not initialized"
}

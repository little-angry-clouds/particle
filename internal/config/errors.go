package config

type particleNotInitialized struct{}

func (e *particleNotInitialized) Error() string {
	return "particle is not initialized"
}

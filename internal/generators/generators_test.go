package generators

import "math/rand"

func init() {
	// Override default RNG generator to create a deterministic one for tests
	newRng = func() *rand.Rand {
		rng := rand.New(rand.NewSource(4))
		return rng
	}
}

package providers

import (
	"os"
	"strings"
)

type EnvironmentProvider struct { }

func (this EnvironmentProvider) GetBytes() ([]byte, error) {

	var out string
	
	for _, e := range(os.Environ()) {
		if ! strings.Contains(e, "\n") && strings.Contains(e, "=") {
			out += e + "\n"
		}
	}
	
	return []byte(out), nil
}

func NewEnvironmentProvider() EnvironmentProvider {
	return EnvironmentProvider{ }
}

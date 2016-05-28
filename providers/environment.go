package providers

import (
	"os"
	"strings"
)

type EnvironmentProvider struct { }

func (this EnvironmentProvider) GetBytes() []byte {

	var out string
	
	for _, e := range(os.Environ()) {
		if ! strings.Contains(e, "\n") && strings.Contains(e, "=") {
			out += e + "\n"
		}
	}
	
	return []byte(out)
}

func NewEnvironmentProvider() EnvironmentProvider {
	return EnvironmentProvider{ }
}

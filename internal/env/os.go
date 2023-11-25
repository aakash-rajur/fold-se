package env

import (
	"os"
	"strings"
)

func loadFromOs() Env {
	env := make(Env)

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		env[pair[0]] = pair[1]
	}

	return env
}

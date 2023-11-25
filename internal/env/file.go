package env

import (
	"github.com/joho/godotenv"
	"log/slog"
	"maps"
	"os"
	"path"
)

func loadFromFile(workDir string, initial Env) Env {
	env := initial

	envFile := path.Join(workDir, ".env")

	_, err := os.Stat(envFile)

	if err != nil {
		if !os.IsNotExist(err) {
			slog.Warn("local environment file not found")
		}

		return env
	}

	dotEnv, err := godotenv.Read(envFile)

	if err != nil {
		slog.Warn("failed to read environment file")

		return env
	}

	if dotEnv == nil {
		return env
	}

	maps.Copy(env, dotEnv)

	return env
}

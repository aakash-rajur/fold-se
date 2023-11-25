package env

func Load(workdir string) Env {
	env := loadFromOs()

	env = loadFromFile(workdir, env)

	return env
}

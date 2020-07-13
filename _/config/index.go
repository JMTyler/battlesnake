package config

import (
	"os"
)

func Initialise() {
	// .argv()
	// .env()
	// .file(`./${process.env.NODE_ENV}.js`, { silent: true })
	// .file('./default.js')
}

func Get(path string, def interface{}) interface{} {
	val, ok := os.LookupEnv(path)
	if !ok {
		return def
	}
	return val
}

package config

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var config = make(map[string]string)

func Get(path string, def string) string {
	initConfig()

	val, ok := config[path]
	if !ok {
		return def
	}
	return val
}

var initialised = false

func initConfig() {
	if initialised {
		return
	}

	defaults, ok := loadYAMLFile("default.yaml")
	if ok {
		set(defaults)
	}

	env, ok := os.LookupEnv("ENVIRONMENT")
	if ok {
		envFile, ok := loadYAMLFile(env + ".yaml")
		if ok {
			set(envFile)
		}
	}

	vars := make(map[string]interface{})
	envvars := os.Environ()
	for _, keypair := range envvars {
		result := strings.Split(keypair, "=")
		vars[result[0]] = result[1]
	}
	set(vars)

	args := make(map[string]interface{})
	for _, keypair := range os.Args[1:] {
		result := strings.Split(keypair, "=")
		args[result[0]] = result[1]
	}
	set(args)

	initialised = true
}

func set(values map[string]interface{}) {
	for key, val := range values {
		config[strings.ToLower(key)] = fmt.Sprintf("%v", val)
	}
}

func loadYAMLFile(filename string) (map[string]interface{}, bool) {
	file, err := os.Open(currentDir() + "/" + filename)
	if err != nil {
		if os.IsNotExist(err) {
			return nil, false
		}
		panic(err)
	}

	var data map[string]interface{}
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&data); err != nil {
		panic(err)
	}
	return data, true
}

func currentDir() string {
	_, filename, _, _ := runtime.Caller(0)
	return filepath.Dir(filename)
}

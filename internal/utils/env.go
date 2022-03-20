package utils

import (
	"fmt"
	"os"
	"strings"
)

func MustGetenv(key string) string {
	env, exist := os.LookupEnv(key)
	if !exist {
		panic(fmt.Sprintf("Env %s is not set, all environs: \n %v", key, strings.Join(os.Environ(), "\n")))
	}
	return env
}

func GetStringEnv(key string, defaultVal string) string {
	env, exist := os.LookupEnv(key)
	env = strings.Trim(env, " ")
	if !exist || env == "" {
		return defaultVal
	}
	return env
}

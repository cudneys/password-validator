package config

import (
	"fmt"
	"os"
	"strconv"
)

func GetEntropyPassword() string {
	return "Numuwh9Mz6RcuLtEBKugqDFqnBhpcQL@PYDz7z6P5qFvBKKTgTtnTycdk*HaYUkNdFcErP@!i^&qvfkBB$4XVSyjRpKD^hSz83QaT!^5PidAEdisvURZqxcrMi2Mgc9k"
}

func GetEnvVarWithDefault(key, defaultValue string) string {
	val := defaultValue
	if v := os.Getenv(key); v != "" {
		val = v
	}
	return val
}

func GetReleaseMode() string {
	return GetEnvVarWithDefault("RELEASE_MODE", "debug")
}

func GetHost() string {
	return GetEnvVarWithDefault("BIND_HOST", "0.0.0.0")
}

func GetPort() string {
	return GetEnvVarWithDefault("BIND_PORT", "8080")
}

func GetMinLen() int {
	val := GetEnvVarWithDefault("PASSWORD_MIN_LENGTH", "16")
	iVal, err := strconv.Atoi(val)
	if err != nil {
		iVal = 16
	}
	return iVal
}

func GetMaxLen() int {
	val := GetEnvVarWithDefault("PASSWORD_MAX_LENGTH", "512")
	iVal, err := strconv.Atoi(val)
	if err != nil {
		iVal = 512
	}
	return iVal
}

func GetBindHost() string {
	port := GetPort()
	host := GetHost()
	return fmt.Sprintf("%s:%s", host, port)
}

func GetLogColorEnabled() bool {
	enabled := true
	if v := os.Getenv("DISABLE_LOGGING_COLOR"); v != "" {
		enabled = false
	}
	return enabled
}

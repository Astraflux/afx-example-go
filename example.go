package example

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

// ReverseString will reverse the string.
func ReverseString(input string) (output string, err error) {
	runes := []rune(input)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	output = string(runes)
	return
}

// PrintEnvAndVars will do just that.
func PrintEnvAndVars() {
	fmt.Print("Starting")

	envVars := make(map[string]string)
	for _, env := range os.Environ() {
		index := strings.Index(env, "=")
		envVars[env[0:index]] = env[index+1:]
	}

	for _, v := range os.Args {
		fmt.Printf("Arguments: %s\n", v)
	}

	for n, v := range envVars {
		fmt.Printf("Environment Variable: %s = %s\n", n, v)
	}

	if val, ok := envVars["TIMEOUT"]; ok {
		if dur, err := time.ParseDuration(val); err == nil {
			fmt.Printf("Sleeping for: %s", dur)
			time.Sleep(dur)
		}
	}

	if val, ok := envVars["ERROR"]; ok {
		if v, err := strconv.Atoi(val); err == nil {
			fmt.Printf("Done: Code = %d", v)
			os.Exit(v)
		} else {
			fmt.Print("Done: Bad error code.")
		}
	} else {
		fmt.Print("Done: No error code.")
	}
}

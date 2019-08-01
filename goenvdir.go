package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
)

func collectEnvVars(envPath string) ([]string, error) {
	var envVars []string
	files, err := ioutil.ReadDir(envPath)
	if err != nil {
		return envVars, err
	}

	for _, file := range files {
		name := file.Name()
		path := filepath.Join(envPath, name)
		data, err := ioutil.ReadFile(path)
		if err != nil {
			return envVars, err
		}
		envVars = append(envVars, name+"="+string(data))
	}
	return envVars, nil
}

func setupCmdEnv(program string, envVars []string, overrideEnv bool) *exec.Cmd {
	cmd := exec.Command(program)
	if overrideEnv {
		cmd.Env = envVars
	} else {
		cmd.Env = append(envVars, os.Environ()...)
	}
	return cmd
}

func main() {
	overrideEnv := flag.Bool("i", false, "Override env")
	flag.Parse()

	if flag.NArg() < 2 {
		log.Fatal("Provide envdir and program to execute")
	}
	envPath, program := flag.Args()[0], flag.Args()[1]

	envVars, err := collectEnvVars(envPath)
	if err != nil {
		log.Fatal(err)
	}

	cmd := setupCmdEnv(program, envVars, *overrideEnv)
	out, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(out))
}

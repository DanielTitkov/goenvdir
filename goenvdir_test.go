package main

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

func TestEnvVarsCollection(t *testing.T) {
	testDir := "./test_dir"
	os.Mkdir(testDir, os.ModePerm)
	defer os.RemoveAll(testDir)

	type envFile struct {
		name string
		data string
	}

	envFiles := []envFile{ // files will be sorted alphabetically
		{"BEST", "FOODISPIZZA"},
		{"FOO", "123"},
		{"Mark", "Stall"},
		{"Zest", "IsOnFire"},
	}

	for _, ef := range envFiles {
		filepath := filepath.Join(testDir, ef.name)
		err := ioutil.WriteFile(filepath, []byte(ef.data), 0755)
		if err != nil {
			t.Error(err)
		}
	}

	envVars, err := collectEnvVars(testDir)
	if err != nil {
		t.Error(err)
	}

	for i, ev := range envVars {
		testcase := envFiles[i].name + "=" + envFiles[i].data
		if ev != testcase {
			t.Errorf("Expected: %s, got %s", testcase, ev)
		}
	}
}

func TestCmdEnvWithoutOverride(t *testing.T) {
	envVars := []string{"FOO=BAR", "JOE=123"}
	cmd := setupCmdEnv("env", envVars, false)
	testEnv := append(envVars, os.Environ()...)
	if len(cmd.Env) != len(testEnv) {
		t.Errorf("Expected len of %d, got %d", len(testEnv), len(cmd.Env))
	}
	for i, ev := range cmd.Env {
		if ev != testEnv[i] {
			t.Errorf("Expected: %s, got %s", testEnv[i], ev)
		}
	}
}

func TestCmdEnvWithOverride(t *testing.T) {
	envVars := []string{"FOO=BAR", "JOE=123"}
	cmd := setupCmdEnv("env", envVars, true)
	testEnv := envVars
	if len(cmd.Env) != len(testEnv) {
		t.Errorf("Expected len of %d, got %d", len(testEnv), len(cmd.Env))
	}
	for i, ev := range cmd.Env {
		if ev != testEnv[i] {
			t.Errorf("Expected: %s, got %s", testEnv[i], ev)
		}
	}
}

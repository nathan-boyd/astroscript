package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

// TestEnvWrapper is a wrapper around file and operating system functions for testing
type testEnvWrapper struct{}

// GetWorkingDirectory returns a fake working directory for testing
func (t *testEnvWrapper) GetWorkingDirectory() (wd string, err error) {
	return
}

func Test_RemoveJpgsCmd(t *testing.T) {

	envWrapper := &testEnvWrapper{}
	cmd := NewRemoveJpgsCmd(envWrapper)
	b := bytes.NewBufferString("")
	cmd.SetOut(b)
	cmd.SetArgs([]string{"--dir", "testisawesome"})
	cmd.Execute()
	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}
	if string(out) != "testisawesome" {
		t.Fatalf("expected \"%s\" got \"%s\"", "testisawesome", string(out))
	}
}

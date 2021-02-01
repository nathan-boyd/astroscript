package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func Test_RemoveJpgsCmd(t *testing.T) {
	cmd := NewRemoveJpgsCmd()
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

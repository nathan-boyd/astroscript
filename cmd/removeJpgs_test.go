package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

func TestAstroscript(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RemoveJpgsCmd Suite")
}

// TestEnvWrapper is a wrapper around file and operating system functions for testing
type testEnvWrapper struct{}

// GetWorkingDirectory returns a fake working directory for testing
func (t *testEnvWrapper) GetWorkingDirectory() (wd string, err error) {
	return
}

var _ = Describe("RemoveJpgsCmd", func() {

	DescribeTable("The directory parameter",

		func(directory string, shouldReturnError bool) {

			commandOutput := bytes.NewBufferString("")
			envWrapper := &testEnvWrapper{}

			cmd := NewRemoveJpgsCmd(envWrapper)
			cmd.SetOut(commandOutput)

			cmd.SetArgs([]string{"--dir", directory})
			cmd.Execute()

			out, err := ioutil.ReadAll(commandOutput)

			Expect(err).Should(BeNil())
			Expect(out).Should(BeEquivalentTo(directory))
		},

		Entry("a string with no context", "foo", false),
	)
})

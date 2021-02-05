package cmd

import (
	"bytes"
	"errors"
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

type testCase struct {
	inputPath      string
	expectedOutput string
	expectedError  error
}

var _ = Describe("The directory parameter", func() {

	DescribeTable("Should be validated", func(testCase testCase) {

		commandOutput := bytes.NewBufferString("")

		cmd := NewRemoveJpgCmd()
		cmd.SetOut(commandOutput)

		cmd.SetArgs([]string{"--dir", testCase.inputPath})
		err := cmd.Execute()

		out, outputErr := ioutil.ReadAll(commandOutput)
		Expect(outputErr).Should(BeNil())

		if nil != testCase.expectedError {
			Expect(err).Should(MatchError(testCase.expectedError))
		} else {
			Expect(err).Should(BeNil())
		}

		if 0 != len(testCase.expectedOutput) {
			Expect(string(out)).Should(ContainSubstring(testCase.expectedOutput))
		}
	},

		Entry("should not return error when input directory exists", testCase{
			inputPath: "/Light/ShouldExist",
		}),

		Entry("should return error when input directory does not exist", testCase{
			inputPath:      "DoesNotExist",
			expectedError:  errors.New("directory does not exist DoesNotExist"),
			expectedOutput: "Usage",
		}),
	)
})

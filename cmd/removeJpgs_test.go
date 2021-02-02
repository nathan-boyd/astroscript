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

// TestEnvWrapper is a wrapper around file and operating system functions for testing
type testEnvWrapper struct {
	DoesDirectoryExist       bool
	WorkingDirectoryVal      string
	WorkingDirectoryErrorVal error
}

// DirectoryExists returns true if directory exists
func (t *testEnvWrapper) DirectoryExists(path string) (directoryExists bool) {
	return t.DoesDirectoryExist
}

// GetWorkingDirectory returns a fake working directory for testing
func (t *testEnvWrapper) GetWorkingDirectory() (wd string, err error) {
	return t.WorkingDirectoryVal, t.WorkingDirectoryErrorVal
}

// GetFilePathSeperator returns the file path seperator for the operating system
func (t *testEnvWrapper) GetFilePathSeperator() (seperator rune) {
	return '/'
}

type testCase struct {
	testEnvWrapper testEnvWrapper
	inputPath      string
	expectedOutput string
	expectedError  error
}

var _ = Describe("The directory parameter", func() {

	DescribeTable("Should be validated", func(testCase testCase) {

		commandOutput := bytes.NewBufferString("")

		cmd := NewRemoveJpgsCmd(&testCase.testEnvWrapper)
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
			testEnvWrapper: testEnvWrapper{
				DoesDirectoryExist: true,
			},
		}),

		Entry("should return error when input directory does not exist", testCase{
			inputPath: "DoesNotExist",
			testEnvWrapper: testEnvWrapper{
				DoesDirectoryExist: false,
			},
			expectedError:  errors.New("directory does not exist DoesNotExist"),
			expectedOutput: "Usage",
		}),

		Entry("should return error when input directory does not contain requires subdirectory", testCase{
			inputPath: "/foo/ShouldExist",
			testEnvWrapper: testEnvWrapper{
				DoesDirectoryExist: true,
			},
			expectedError: errors.New(SubdirectoryNotfoundMessage),
		}),
	)
})

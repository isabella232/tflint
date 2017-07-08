package printer

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/wata727/tflint/issue"
)

func TestDefaultPrint(t *testing.T) {
	cases := []struct {
		Name   string
		Input  []*issue.Issue
		Result string
	}{
		{
			Name:   "no issues",
			Input:  []*issue.Issue{},
			Result: successColor("Awesome! Your code is following the best practices :)") + "\n",
		},
		{
			Name: "multi files",
			Input: []*issue.Issue{
				{
					Detector: "error detector",
					File:     "template.tf",
					Line:     1,
					Type:     "ERROR",
					Message:  "example error message",
				},
				{
					Detector: "notice detector",
					File:     "application.tf",
					Line:     10,
					Type:     "NOTICE",
					Message:  "example notice message",
					Link:     "https://github.com/wata727/tflint",
				},
				{
					Detector: "warning detector",
					File:     "template.tf",
					Line:     5,
					Type:     "WARNING",
					Message:  "example warning message",
				},
				{
					Detector: "warning detector",
					File:     "template.tf",
					Line:     3,
					Type:     "WARNING",
					Message:  "example warning message",
				},
			},
			Result: fmt.Sprintf(`%s
	%s example notice message (notice detector)
%s
	%s example error message (error detector)
	%s example warning message (warning detector)
	%s example warning message (warning detector)

Result: %s  (%s , %s , %s)
`, fileColor("application.tf"), noticeColor("NOTICE:10"), fileColor("template.tf"), errorColor("ERROR:1"), warningColor("WARNING:3"), warningColor("WARNING:5"), fileColor("4 issues"), errorColor("1 errors"), warningColor("2 warnings"), noticeColor("1 notices")),
		},
	}

	for _, tc := range cases {
		stdout := &bytes.Buffer{}
		stderr := &bytes.Buffer{}
		p := NewPrinter(stdout, stderr)
		p.DefaultPrint(tc.Input)
		result := stdout.String()

		if result != tc.Result {
			t.Fatalf("\nBad: %s\nExpected: %s\n\ntestcase: %s", result, tc.Result, tc.Name)
		}
	}
}

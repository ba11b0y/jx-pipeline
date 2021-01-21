package processor_test

import (
	"github.com/jenkins-x/jx-pipeline/pkg/pipelines/processor"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConvertLegacyStepAnnotationURLToUsesImage(t *testing.T) {
	testCases := []struct {
		text     string
		expected string
	}{
		{
			text:     "https://raw.githubusercontent.com/jenkins-x/jx3-pipeline-catalog/60bed6408732c1eda91a15713f51a9f97dcb1757/tasks/git-clone/git-clone-pr.yaml",
			expected: "uses:jenkins-x/jx3-pipeline-catalog/tasks/git-clone/git-clone-pr.yaml@60bed6408732c1eda91a15713f51a9f97dcb1757",
		},
		{
			text:     "https://something/cheese.yaml",
			expected: "uses:https://something/cheese.yaml",
		},
		{
			text:     "",
			expected: "",
		},
	}
	for _, tc := range testCases {
		actual := processor.ConvertLegacyStepAnnotationURLToUsesImage(nil, "")
		assert.Equal(t, tc.expected, actual, "for annotation value %s", tc.text)
		t.Logf("converted %s => %s\n", tc.text, actual)
	}
}

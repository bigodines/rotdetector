package rotdetector

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseFileCanLoadFile(t *testing.T) {
	testFilePath := "fixtures/test_file1.js"
	// Ensure the test file exists
	if _, err := os.Stat(testFilePath); os.IsNotExist(err) {
		t.Fatalf("Test file does not exist: %s", testFilePath)
	}

	// Test with TODO and verbose flags set to true
	opts := ParseOptions{Path: testFilePath, Todo: true, Verbose: true}
	detected, err := ParseFile(opts)
	if err != nil {
		t.Fatalf("ParseFile returned an error: %v", err)
	}
	assert.Equal(t, detected, true, "Expected a detected rot")

	// Test with TODO and verbose flags set to false
	opts = ParseOptions{Path: testFilePath, Todo: false, Verbose: false}
	_, err = ParseFile(opts)
	if err != nil {
		t.Fatalf("ParseFile returned an error: %v", err)
	}
}

func TestDetectLanguage(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{"main.go", "golang"},
		{"script.ts", "javascript"},
		{"script.js", "javascript"},
		{"script.rb", "ruby"},
		{"script.py", "python"},
		{"unknown.txt", ""},
	}

	for _, tt := range tests {
		t.Run(tt.path, func(t *testing.T) {
			result := detectLanguage(tt.path)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestGetCommentRegex(t *testing.T) {
	tests := []struct {
		language string
		content  string
		expected bool
	}{
		{"golang", "// This is a comment", true},
		{"golang", "func Thisisacomment()", false},
		{"javascript", "// This is a comment", true},
		{"ruby", "# This is a comment", true},
		{"ruby", "// This isn't a comment!!!", false},
		{"python", "# This is a comment", true},
		{"unknown", "// This should not match", false},
	}

	for _, tt := range tests {
		t.Run(tt.language, func(t *testing.T) {
			regex := getCommentRegex(tt.language)
			if regex == nil {
				if tt.expected {
					t.Errorf("expected a valid regex, got nil")
				}
				return
			}
			match := regex.MatchString(tt.content)
			if match != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, match)
			}
		})
	}
}

func TestParseContent(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		content  string
		language string
		todo     bool
		verbose  bool
		expected bool
	}{
		{
			name:     "Expired BestBy",
			path:     "test.go",
			content:  "// BestBy: 01/2001: foobar\n",
			language: "golang",
			todo:     false,
			verbose:  true,
			expected: true,
		},
		{
			name:     "Valid BestBy",
			path:     "test.go",
			content:  "// BestBy: 01/2099\n",
			language: "golang",
			todo:     false,
			verbose:  false,
			expected: false,
		},
		{
			name:     "TODO Comment",
			path:     "test.go",
			content:  "// TODO: fix this\n",
			language: "golang",
			todo:     true,
			verbose:  false,
			expected: false,
		},
		{
			name:     "No Comments",
			path:     "test.go",
			content:  "func main() {}\n",
			language: "golang",
			todo:     false,
			verbose:  false,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			result := parseContent(tt.path, []byte(tt.content), tt.language, tt.todo, tt.verbose)
			if result != tt.expected {
				t.Errorf("expected %v, got %v", tt.expected, result)
			}
		})
	}
}

func TestParseFileErrors(t *testing.T) {
	// Test with a non-existent file
	opts := ParseOptions{Path: "non_existent_file.go", Todo: false, Verbose: false}
	_, err := ParseFile(opts)
	if err == nil {
		t.Fatalf("expected an error for non-existent file, got nil")
	}
}

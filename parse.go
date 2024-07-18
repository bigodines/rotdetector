package rotdetector

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// ParseFile Parses a single file searching for BestBy annotations
// BestBy: 12/2024: will only error if the file cannot be read but will only understand context around predefined
// number of languages
func ParseFile(path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("error reading file %q: %v", path, err)
	}

	language := detectLanguage(path)
	if language != "" {
		parseContent(path, content, "golang")
	}

	return nil
}

// BestBy 01/2001 - another example
func parseContent(path string, content []byte, language string) {
	commentRegex := getCommentRegex(language)
	re := regexp.MustCompile(`BestBy[\s\(\-\:]?\d{2}/\d{4}`)
	lines := strings.Split(string(content), "\n")
	for n, line := range lines {
		matches := commentRegex.FindAllString(line, -1)
		if len(matches) > 0 {
			for _, comment := range matches {
				bestByMatches := re.FindAllString(comment, -1)
				if len(bestByMatches) > 0 {
					l := n
					if n+1 < len(lines) {
						l = n + 1
					} else {
						l = n
					}
					fmt.Printf("File: %s\n (L:%d) %s \n\t-> %v", path, l, comment, lines[l])

				}
			}
		}
	}
}

func detectLanguage(path string) string {
	ext := filepath.Ext(path)
	switch ext {
	case ".go":
		return "golang"
	case ".ts", ".tsx", ".js", ".jsx":
		return "javascript"
	case ".rb":
		return "ruby"
	default:
		return ""
	}
}

func getCommentRegex(language string) *regexp.Regexp {
	switch language {
	case "golang":
		return regexp.MustCompile(`(?m)\/\/.*|\/\*[\s\S]*?\*\/`)
	case "javascript":
		return regexp.MustCompile(`(?m)\/\/.*|\/\*[\s\S]*?\*\/`)
	case "ruby":
		return regexp.MustCompile(`(?m)#.*`)
	default:
		return nil
	}
}

package rotdetector

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type ParseOptions struct {
	Path    string
	Todo    bool
	Verbose bool
}

// ParseFile Parses a single file searching for BestBy annotations
// BestBy: 12/2024: will only error if the file cannot be read but will only understand context around predefined
// number of languages
func ParseFile(opts ParseOptions) error {
	content, err := os.ReadFile(opts.Path)
	if err != nil {
		return fmt.Errorf("error reading file %q: %v", opts.Path, err)
	}

	language := detectLanguage(opts.Path)
	if language != "" {
		parseContent(opts.Path, content, "golang", opts.Todo, opts.Verbose)
	}

	return nil
}

// BestBy 01/2001 - another example
func parseContent(path string, content []byte, language string, todo bool, verbose bool) {
	commentRegex := getCommentRegex(language)
	// I have an interesting problem to solve. Let me use Regex. Now I have two problems to solve.
	reBestBy := regexp.MustCompile(`[bB]est[bB]y[\s\(\-\:]?(?P<Month>\d{1,2})/(?P<Year>\d{2}|\d{4})`)
	reTodo := regexp.MustCompile(`TODO`)
	lines := strings.Split(string(content), "\n")
	for n, line := range lines {
		matches := commentRegex.FindAllString(line, -1)
		if len(matches) > 0 {
			for _, comment := range matches {
				bestByMatches := reBestBy.FindStringSubmatch(comment)
				if len(bestByMatches) > 0 {
					monthStr := bestByMatches[1]
					yearStr := bestByMatches[2]
					month, err := strconv.Atoi(monthStr)
					if err != nil {
						fmt.Errorf("error parsing month: %v", err)
						continue
					}
					year, err := strconv.Atoi(string(yearStr))
					if err != nil {
						fmt.Errorf("error parsing year: %v", err)
						continue
					}
					// Adjust year for 2-digit format
					if year < 100 {
						year += 2000
					}
					bestByDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
					currentDate := time.Now()
					if bestByDate.Before(currentDate) {
						// at this point we've detected an expired comment
						l := n
						if n+1 < len(lines) {
							l = n + 1
						} else {
							l = n
						}
						fmt.Printf(Green+"File: %s\n (L:%d) %s \n\t-> %v\n"+Reset, path, l, comment, lines[l])
					}
				}
				// no need to analyze further if we are not looking for TODOs
				if !todo {
					continue
				}
				// TODO: combine with the block above.
				reTodoMatches := reTodo.FindAllString(comment, -1)
				if len(reTodoMatches) > 0 {
					l := n
					if n+1 < len(lines) {
						l = n + 1
					} else {
						l = n
					}
					fmt.Printf(Yellow+"File: %s\n (L:%d) %s \n\t-> %v\n"+Reset, path, l, comment, lines[l])
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

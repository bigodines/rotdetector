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

const (
	TypeBestBy = iota
	TypeTodo
)

type ParseOptions struct {
	Path    string
	Todo    bool
	Verbose bool
}

type ParseResult struct {
	Line int
	File string
	Type int
}

// ParseFile Parses a single file searching for BestBy annotations
// BestBy: 12/2024: will only error if the file cannot be read but will only understand context around predefined
// number of languages
func ParseFile(opts ParseOptions) (foundRot bool, err error) {
	content, err := os.ReadFile(opts.Path)
	if err != nil {
		return true, fmt.Errorf("error reading file %q: %v", opts.Path, err)
	}

	language := detectLanguage(opts.Path)
	if language != "" {
		foundRot = parseContent(opts.Path, content, "golang", opts.Todo, opts.Verbose)
	}

	return foundRot, nil
}

// BestBy 01/2001 - another example
func parseContent(path string, content []byte, language string, todo bool, verbose bool) bool {
	foundRot := false
	commentRegex := getCommentRegex(language)
	// I have an interesting problem to solve. Let me use Regex. Now I have two problems to solve.
	// This regexp matches several combinations of "BestBy" annotations
	reBestBy := regexp.MustCompile(`[bB]est[bB]y[\s\(\-\:]?(?P<Month>\d{1,2})/(?P<Year>\d{4}|\d{2})`)
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
					// Adjust year from 2-digit format
					if year < 100 {
						year += 2000
					}

					bestByDate := time.Date(year, time.Month(month), 1, 0, 0, 0, 0, time.UTC)
					currentDate := time.Now()
					if bestByDate.Before(currentDate) {
						foundRot = true
						// at this point we've detected an expired comment
						l := n
						if n+1 < len(lines) {
							l = n + 1
						}
						fmt.Printf(Magenta+"File: %s (L:%d) %s \n\t-> %v\n"+Reset, path, l, comment, lines[l])
					} else {
						Debug(fmt.Sprintf("Found BestBy that is still valid: %v", comment))
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
					fmt.Printf(Cyan+"File: %s (L:%d) %s \n\t-> %v\n"+Reset, path, l, comment, lines[l])
				}
			}
		}
	}
	return foundRot
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
	cStyleComment := regexp.MustCompile(`(?m)\/\/.*|\/\*[\s\S]*?\*\/`)

	switch language {
	case "golang":
		return cStyleComment
	case "javascript":
		return cStyleComment
	case "ruby":
		return regexp.MustCompile(`(?m)#.*`)
	default:
		return nil
	}
}

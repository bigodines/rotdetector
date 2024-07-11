package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
)

func ParseFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("error reading file %q: %v\n", path, err)
		return
	}

	// language := detectLanguage(path)
	// if language == "" {
	// 	return
	// }

	parseContent(path, content, "golang")
}

// BestBy 01/2001 - another example
func parseContent(path string, content []byte, language string) {
	commentRegex := getCommentRegex(language)
	re := regexp.MustCompile(`BestBy[\s\(\-\:]?\d{2}/\d{4}`)
	matches := commentRegex.FindAllString(string(content), -1)
	if len(matches) > 0 {
		for _, comment := range matches {
			bestByMatches := re.FindAllString(comment, -1)
			if len(bestByMatches) > 0 {
				fmt.Printf("File: %s\n", path)
				for _, match := range bestByMatches {
					fmt.Println(" ", match)
				}
			} else {
				fmt.Printf("no matches in %s\n", path)
			}
		}
	}
}

func detectLanguage(path string) string {
	ext := filepath.Ext(path)
	println(ext)
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

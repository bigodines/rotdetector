package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"

	"github.com/bigodines/bigopool"
)

type (
	ParseJob struct {
		fileName string
	}
)

func (j ParseJob) Execute() (bigopool.Result, error) {
	println(j.fileName)
	parseFile(j.fileName)
	// your logic here.
	// Result is an interface{}
	return "anything", nil
}

func main() {
	dir := flag.String("dir", ".", "Directory to start parsing from")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	dispatcher, err := bigopool.NewDispatcher(50, 1000)
	if err != nil {
		panic(err)
	}
	// spawn workers
	dispatcher.Run()
	// send work items
	err = filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			job := ParseJob{fileName: path}
			dispatcher.Enqueue((job))
		}
		return nil
	})
	if err != nil {
		log.Fatalf("error walking the path %q: %v\n", *dir, err)
	}
}

func parseFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Printf("error reading file %q: %v\n", path, err)
		return
	}

	language := detectLanguage(path)
	if language == "" {
		return
	}

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

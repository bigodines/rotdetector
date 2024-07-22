package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/bigodines/bigopool"
	rd "github.com/bigodines/rotdetector"
)

type (
	ParseJob struct {
		fileName string
		todo     bool
	}
)

// Parses a single file searching for BestBy annotations
// compares the date with the current date and flags the file if the date is in the past
func (j ParseJob) Execute() (bigopool.Result, error) {
	opts := rd.ParseOptions{Path: j.fileName, Todo: j.todo}
	return rd.ParseFile(opts)
}

func main() {
	// Configure
	dir := flag.String("dir", ".", "Directory to start parsing from")
	v := flag.Bool("v", false, "Verbose (debug) mode")
	ci := flag.Bool("ci", false, "(soon) CI friendly mode (no color output, exit 1 when detect rot)")
	todo := flag.Bool("todo", false, "detect TODOs")
	export := flag.String("export", "", "Export results to a file")
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options]\n", os.Args[0])
		fmt.Fprintf(flag.CommandLine.Output(), "Options:\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	total := 0

	if v != nil && *v {
		rd.SetLogLevel(rd.DEBUG)
		rd.Debug("Now running in debug mode")
	}

	if *ci {
		// disable colors
		rd.Cyan = ""
		rd.Red = ""
		rd.Magenta = ""
		rd.Reset = ""
		rd.Debug("-ci mode is set")
	}

	if *export != "" {
		rd.Debug("Exporting results to: ", *export)
	}

	if *todo {
		rd.Debug("Detecting TODOs")
	}

	err := filepath.Walk(*dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			job := ParseJob{fileName: path, todo: *todo}
			job.Execute()
			total++
		}
		return nil
	})
	if err != nil {
		log.Fatalf("error walking the path %q: %v\n", *dir, err)
	}
	// rotten, errs := dispatcher.Wait()
	// if len(errs.All()) > 0 {
	// 	log.Fatalf("errors: %v\n", errs)
	// }

	rd.Debug(fmt.Sprintf("Done scanning. Total files scanned: %d", total))
	// if len(rotten) > 0 {
	// 	for _, r := range rotten {
	// 		if r.(bool) {
	// 			os.Exit(1)
	// 		}

	// 	}
	// }

}

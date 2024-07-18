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
	// BestBy: 02/2023 - This is just an example of how to use RotDetector
	ParseJob struct {
		fileName string
	}
)

// Parses a single file searching for BestBy annotations
// compares the date with the current date and flags the file if the date is in the past
func (j ParseJob) Execute() (bigopool.Result, error) {
	rd.ParseFile(j.fileName)
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

	dispatcher, err := bigopool.NewDispatcher(32, 1000)
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
	_, errs := dispatcher.Wait()
	if len(errs.All()) > 0 {
		log.Fatalf("errors: %v\n", errs)
	}
}
# RotDetector

This amazing piece of technology has been built to assist detection of code that is rotting. Hook it up to your CI/CD or build tool and simply add `BestBy MM/YYYY` to a comment. RotDetector will parse (very quickly) all the files looking for expired notes and alert on those.

## Quickstart

```bash
git clone git@github.com:bigodines/rotdetector.git
cd rotdetector
make build
./bin/rotdetector -dir=.
```

These commands will download, build and run the latest version of rotdetector in the diven directory `-dir` scanning all the files that contain a line comment `BestBy MM/YYYY` or [optionally] `TODO` on it (and its subdirectories). It will print the file, comment, line number and the immediate line below the comment which will likely identify what needs out attention.

Advanced options include custom outputs to make it easier to plug into your preferred CI/CD pipeline

## Options

```bash
Usage: ./bin/rotdetector [options]
Options:
  -ci
    	(soon) CI friendly mode (no color output, exit 1 when detect rot)
  -dir string
    	Directory to start parsing from (default ".")
  -export string
    	Export results to a file
  -todo
    	detect TODOs
  -v	Verbose (debug) mode
```
# rotDetector

This amazing piece of technology has been built to assist detecting code that is rotting. Hook it up to your CI/CD or build tool and simply add `BestBy MM/YYYY` to a comment. RotDetector will parse (very quickly) your files looking for expired notes and alert on those.

You can then decide if you want to deal with the issue or change the date and document the reasoning in a pull-request.

## Quickstart

```bash
git clone git@github.com:bigodines/rotdetector.git
cd rotdetector
make build
./bin/rotdetector -dir=.
```

These commands will download, build and run the latest version of rotdetector in the given directory `-dir` (and its subdirectories) scanning all the files that contain a line comment `BestBy MM/YYYY` on it. It will print the file, comment, line number and the immediate line below the comment which will likely identify what needs your attention. It'll also exit(1)

## Options

```bash
Usage: ./bin/rotdetector [options]
Options:
  -ci
    	CI friendly mode (no color output)
  -dir string
    	Directory to start parsing from (default ".")
  -export string
    	(soon) Export results to a file
  -todo
    	detect TODOs
  -v	Verbose (debug) mode
```

## Learn more about rotDetector

https://imbigo.net/posts/rotdector/

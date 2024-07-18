# RotDetector

This amazing piece of technology has been built to assist detection of code that is rotting. Hook it up to your CI/CD or build tool and simply add `BestBy MM/YYYY` to a comment. RotDetector will parse (very quickly) all the files looking for expired notes and alert on those.

## Quickstart

```bash
git clone git@github.com:bigodines/rotdetector.git
cd rotdetector
make build
./bin/rotdetector -dir=.
```

These commands will download, build and run the latest version of rotdetector in the current directory


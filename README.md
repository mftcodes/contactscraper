# Contact Scraper
## Purpose
Search provided websites for contact information and return in .csv form. The finished idea will read from `/srv/scraper.in/` and write to `/srv/scraper.out/`. When run in `debug mode` the program reads from and writes too `/tmp/scraper.in/` and `/tmp/scraper.out/` respectively.

This project is meant to show some of the skills within Go (golang) that I have been building on a closed source project, as well as continue to practice and sharpen those skills, rather than being a optimzed scraper for general use. 

## Usage
Once installed or compiled, the softare is as simply as calling one command from the command line:
```
scraper {flag}
```
### Flags
* `-d`: Sets the program into debug mode, and will read/write from the linux based file system at `/tmp/scraper.in/` and `/tmp/scraper.out/`

## Setup local dev
1. From project root, first initialize with
    ```
    go mod init scraper
    go mod tidy
    ```
2. Now traverse the project file system and build each module (e.g. `helpers`, `httpservice`, `readservice`, `writeservice`)
    1. CD into each directory, and build, e.g.:
        ```
        cd helpers
        go build .
        ```
    2. Repeat this for each directory that contains go files
3. Return to the root, and install
    ```
    go install scraper
    ```
4. You can now run it simply by using the shell command `scraper`
5. If the command is not found you need to export PATH for the newly installed program with
    ```
    export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
    ```
6. For `Debug Mode` use the optional flag  `-d`
8. Optional, when you are done developin and want to clear modcache run `go clean -modcache`

### Notes on current state
Reads URLs from a single file then retrieves those pages html and saves it to a file. Fruther processing and navigating to common contacts pages through pattern matching is being worked on. 
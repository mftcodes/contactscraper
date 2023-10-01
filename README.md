# contactscraper
## Purpose
Search provided websites for contact information and return in .csv form. 

## Setup local dev
1. From project root run the following
    ```
    go mod init scraper
    go mod tidy
    cd helpers
    go build .
    cd ..
    go mod install scarper
    ```
2. Export PATH for newly installed program with
    ```
    export PATH=$PATH:$(dirname $(go list -f '{{.Target}}' .))
    ```
3. Now you should be able to run with the command `scraper`.
4. For `Debug Mode` use the optional flag  `-d`
5. Optional, if you want to clear modcache run `go clean -modcache`

### Notes
Just getting started. Work in progress.
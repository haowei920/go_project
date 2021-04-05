# file-search-api

file-search-api provides a tool for users to recursively find all filenames (satisfying the regex) in a select directory within the container.

## Installation
Ensure that you have docker installed on your local machine. If not, please download it from the official docker website for your local machine's OS

Follow the steps below to install the docker image and run the container in interactive mode.

```bash
docker pull haowei920/file-search-api
docker run -it --entrypoint /bin/bash haowei920/file-search-api

```

## Usage

Follow the step below to recursively find all filenames (satisfying the regex) in a select directory within the container. 

```bash
cd go_project/server
./server.sh &
cd /build/go_project/client
./cli.sh -name=<file_name_regex> -path=<absolute path name of directory to search>
```

If the program manages to find files within directory that satisfies the regex, it would print the absolute path of all such files after below the line executing the script. Each result is separated by a new line. Example:
```bash
./cli.sh -name=cli.* -path=/build/go_project/client/
/build/go_project/client/cli
/build/go_project/client/cli.go
/build/go_project/client/cli.sh
```
If the directory path provided is invalid, it would return the following
```bash
./cli.sh -name=cli.* -path=/build/go_project/gibberish/
This path cannot be found
```
If the directory path provided is valid, but no files within the directory satisfies the regex provided, nothing will be printed



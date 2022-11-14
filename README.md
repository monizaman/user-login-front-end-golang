# User Management Front End

**Setup Guide**

## Prerequisites

- Download and install go: https://golang.org/dl/

## Setup configuration

- Copy `config.yml.sample` file from `root` directory and create `config.yml` in the same directory
- Open file `config.yml` and put appropriates

## Install dependencies and Run the App Locally

*Follow the steps given bellow to build and run this project:*

```bash
# Create Go workspace and move project to the workspace

# Install dependencies (only first time)
go mod tidy

# Run project:
go run main.go
```

## Run using Docker

- Copy `config.yml.sample` file from `root` directory and create `config.yml` in the same directory
- Open file `config.yml` and put appropriates
- Run `bash build-run.sh`

### Important Go Commands

- Get GOPATH: `go env GOPATH`
- Set GOBIN directory (Put your URL)
    - Mac/Ubuntu: `export GOBIN=/Users/name/go/bin`
    - Windows: `SET GOBIN=/Users/name/go/bin`
- Check Go environment variables: `go env`
- Update packages to avoid security and future compatibility
    - Step 1. `go get -u all`
    - Step 2. `go mod tidy`

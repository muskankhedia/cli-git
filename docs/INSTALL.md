# :minidisc: Installation Instructions

## Prerequisites

To run this project, the system should have the following prerequisites, if not then click on the link mentioned in the prerequites and install it and then follow the installation steps.

The prerequisites for installing the project :
1. [golang](https://golang.org/dl/ "Install GOLang")


# Installation Steps
1. Clone the repository in your GOPATH

    ```bash
        $ go get https://github.com/muskankhedia/cli-git
        $ cd $GOPATH/src/github.com/muskankhedia/cli-git
    ```

## Steps to build the tool

2. Install the cli-tool to the local system

    ```
        $ go build -o $GOPATH/bin/cli-git
    ```
3. Use the terminal in any location of the local system

    ```bash
        $ cli-git
    ```

## Steps to run the project for development

2. Run the project 

    ```
        go run main.go <command>
    ```


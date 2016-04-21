# golang-samples
Playground to test and learn go fundamentals

## Workspace layout

Go code is kept in a workspace that contains multiple source repositories (e.g. git). If you follow the Go conventions, the Go tool understands the layout of a workspace. You don't need Makefiles. A workspace has this layout:

    $GOPATH/
        src/
            github.com/kspichale/golang-samples/
                types
                    circle.go
        bin/
           circle
           
## Create a workspace

Create the workspace directory and set the environment variable GOATH.

    mkdir /foo/bar/workspace
    export GOPATH=/foo/bar/workspace
   
The GOPATH environment variable tells the Go tool where can find the workspace.

    go get github.com/kspichale/golang-samples
   
This command fetches the source repository from github and places it into the workspace. The Go tool recognizes "github.com" in the package path.

    go install github.com/kspichale/golang-samples/types
   
Finally the Go install command is used to build the sources. The binary is placed in $GOPATH/bin/.
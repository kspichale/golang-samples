A playground to test and learn Go fundamentals.

## Workspace layout

Go code is kept in a workspace that contains multiple source repositories (e.g. git). If you follow the Go conventions, the Go tool understands the layout of a workspace. You don't need Makefiles. Everyone in the community should use the same layout to share code easily.

This workspace has this layout:

    $GOPATH/
        src/
            github.com/kspichale/golang-samples/
                types
                    circle.go
        bin/
           circle

A more general Go workspace could follow this layout pattern:

    $GOPATH/
        src/
            myproj/
                mypack
                    lib.go
                    lib_test.go
            main/
                myapp.go
        bin/
           myapp

## Create a workspace

Create the workspace directory and set the environment variable GOPATH.

    mkdir /foo/bar/workspace
    export GOPATH=/foo/bar/workspace

The GOPATH environment variable tells the Go tool where it can find the workspace.

    go get github.com/kspichale/golang-samples

This command fetches the source repository from github and places it into the workspace. The Go tool recognizes "github.com" in the package path.

    go install github.com/kspichale/golang-samples/types

Finally the Go install command is used to build the sources. The binary is placed in $GOPATH/bin/. What else can you do with to Go tool?

    build     compile pkg and deps
    get       download and install pkg and deps
    install   compile and install pkg and deps
    test      test pkg

You now can continue fetching other repositories. You could for instance install the testing library [gocheck](https://labix.org/gocheck) that is used in circle_test.go.

    go get gopkg.in/check.v1

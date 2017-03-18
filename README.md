# BuildInfo
Provide basic build information


## Usage
In shell of your choice zsh/bash/sh,  
cd to your project

Put project path in GOPATH

        export GOPATH=$(PWD):$(go env GOPATH)
        go get github.com/naqernf/buildinfo

Set var

        importname=github.com/naqernf/buildinfo
        buildstamp=$(date '+%s')
        gitdescribe=$(git describe --always --tag --dirty)
        githash=$(git rev-parse HEAD)
        build_ldflags="-X ${importname}.name=${gitdescribe} -X ${importname}.buildstamp=${buildstamp} -X ${importname}.githash=${githash}"

Build example

        go build -ldflags="${build_ldflags}" -o ./bin/example-${gitdescribe} github.com/naqernf/buildinfo/example

Run example

        ./bin/example-${gitdescribe}


Feedback is highly appreciated!  
Please create an issue, if something is not working for you.

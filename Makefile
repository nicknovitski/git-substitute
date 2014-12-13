all: test

deps:
	go get github.com/docopt/docopt-go
	go get github.com/mitchellh/gox

dist: deps
	gox -os="darwin linux" -arch="386 amd64" -output="{{.OS}}/{{.Arch}}/git-substitute" -verbose

bats :
	git clone https://github.com/sstephenson/bats.git

git-substitute : deps
	go build

test: bats git-substitute
	go test
	bats/bin/bats test

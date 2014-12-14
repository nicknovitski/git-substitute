default: test

deps:
	go get github.com/docopt/docopt-go
	go get github.com/mitchellh/gox

dist: deps
	gox -os="darwin linux" -arch="386 amd64" -output="{{.OS}}/{{.Arch}}/git-substitute" -verbose

bats :
	git clone https://github.com/sstephenson/bats.git

bin/git-substitute: deps
	go build -o bin/git-substitute

test: bats bin/git-substitute
	bats/bin/bats test

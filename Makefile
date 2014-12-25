default: test

deps:
	go get github.com/docopt/docopt-go

dist: deps
	go get github.com/mitchellh/gox
	gox -os="darwin linux" -arch="386 amd64" -output="dist/{{.OS}}/{{.Arch}}/git-substitute" -verbose
	tar --create --gzip --verbose --file=dist/git_substitute_linux_amd64.tar.gz --directory=dist/linux/amd64 git-substitute
	tar --create --gzip --verbose --file=dist/git_substitute_linux_386.tar.gz --directory=dist/linux/386 git-substitute
	tar --create --gzip --verbose --file=dist/git_substitute_darwin_amd64.tar.gz --directory=dist/darwin/amd64 git-substitute
	tar --create --gzip --verbose --file=dist/git_substitute_darwin_386.tar.gz --directory=dist/darwin/386 git-substitute

bats :
	git clone https://github.com/sstephenson/bats.git

bin/git-substitute: deps
	go build -o bin/git-substitute

test: bats bin/git-substitute
	bats/bin/bats test

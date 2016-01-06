default: test

dist/git_substitute_linux_386.tar.gz :
	stack build --install-ghc --os linux --arch i386 --local-bin-path ./bin --copy-bins
	tar --create --gzip --verbose --file=dist/git_substitute_linux_386.tar.gz --directory=bin/ git-substitute

dist/git_substitute_linux_amd64.tar.gz :
	stack build --install-ghc --os linux --arch x86_64 --local-bin-path ./bin --copy-bins
	tar --create --gzip --verbose --file=dist/git_substitute_linux_amd64.tar.gz --directory=bin git-substitute

dist/git_substitute_darwin_386.tar.gz :
	stack build --install-ghc --os osx --arch i386 --local-bin-path ./bin --copy-bins
	tar --create --gzip --verbose --file=dist/git_substitute_darwin_386.tar.gz --directory=bin git-substitute

dist/git_substitute_darwin_amd64.tar.gz :
	stack build --install-ghc --os osx --arch x86_64 --local-bin-path ./bin --copy-bins
	tar --create --gzip --verbose --file=dist/git_substitute_darwin_amd64.tar.gz --directory=bin git-substitute

dist: dist/git_substitute_linux_amd64.tar.gz dist/git_substitute_darwin_amd64.tar.gz dist/git_substitute_linux_386.tar.gz dist/git_substitute_darwin_386.tar.gz

bats :
	git clone https://github.com/sstephenson/bats.git

bin/git-substitute: git-substitute.cabal stack.yaml $(shell ls **/*.hs)
	stack build --no-terminal --install-ghc --local-bin-path ./bin --copy-bins

test: bats bin/git-substitute
	bats/bin/bats test

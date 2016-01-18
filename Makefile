default : git-substitute.cabal stack.yaml $(shell ls **/*.hs)
	stack build --no-terminal --install-ghc --local-bin-path ./bin --copy-bins

docker : git-substitute.cabal stack.yaml $(shell ls **/*.hs)
	stack build --no-terminal --install-ghc --local-bin-path ./bin --copy-bins --docker

get-deps : bats
	stack docker pull
	stack setup

bats :
	git clone https://github.com/sstephenson/bats.git

test : bats default
	bats/bin/bats test

test-docker : bats docker
	bats/bin/bats test

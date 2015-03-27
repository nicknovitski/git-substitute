FROM 1.4.2-cross

RUN go get github.com/docopt/docopt-go

RUN git clone https://github.com/sstephenson/bats.git

CMD go build -o bin/git-substitute && bats/bin/bats test 

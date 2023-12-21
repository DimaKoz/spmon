

.PHONY: mon
mon:
	tput bold setaf 2;echo "building monitor";tput sgr0
	go build -o ./cmd/monitor/monitor ./cmd/monitor/*.go

.PHONY: clean
clean:
	rm -f ./cmd/monitor/monitor
	rm -f ./log*.txt
	rm -f ./tempfile*
	rm -f ./cover.html cover.out coverage.txt

.PHONY: lnt
lnt:
# excluded 'paralleltest' by the reason - not now
# excluded 'wsl' by the reason - 'wsl' and 'gofumpt' fights between each other
# excluded 'gochecknoglobals' by reason - I need global variables sometimes
# excluded 'exhaustivestruct' - deprecated
# excluded 'depguard' - no need in it
	tput bold setaf 1;golangci-lint run --version;tput sgr0
# golangci-lint run -v --enable-all --disable gochecknoglobals --disable paralleltest --disable exhaustivestruct --disable depguard --disable wsl
	golangci-lint run -v

.PHONY: fmt
fmt:
	# to install it:
	# go install mvdan.cc/gofumpt@latest
	gofumpt -l -w .

.PHONY: gci
gci:
	# to install it:
	# go install github.com/daixiang0/gci@latest
	# after that add a place of binaries to $PATH
	# export PATH=$PATH:/your path/go/bin
	gci write --skip-generated -s default .

.PHONY: gofmt
gofmt:
	gofmt -s -w .

.PHONY: fix
fix: gofmt gci fmt


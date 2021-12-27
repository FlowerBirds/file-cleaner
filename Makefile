.PHONY: build

VERSION=0.0.3
BANIRY=file-cleaner.exe

build:
	@go build -ldflags "-X 'main.Version=${VERSION}'" -o ${BANIRY} main.go
	@echo output binary file: ${BANIRY} version ${VERSION}
	@echo buil success
	@${BANIRY}
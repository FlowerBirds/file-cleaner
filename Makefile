.PHONY: build

VERSION=0.0.7
BANIRY=file-cleaner

GO111MODULE=on

buildcmd = GOOS=$(1) GOARCH=$(2) go build -ldflags "-X 'main.Version=${VERSION}'" -o build/$(BANIRY)$(3)
tar = cd build && tar -cvzf ${BANIRY}-$(1)_$(2)-${VERSION}.tar.gz $(BANIRY)$(3) && rm $(BANIRY)$(3)
test = cd build && tar -xvzf ${BANIRY}-$(1)_$(2)-${VERSION}.tar.gz


build: build/linux_amd64 build/linux_arm64 build/windows_amd64 test/windows

build/linux_amd64:
	$(call buildcmd,linux,amd64,)
	$(call tar,linux,amd64)

build/linux_arm64:
	$(call buildcmd,linux,arm64,)
	$(call tar,linux,arm64)

build/windows_amd64:
	$(call buildcmd,windows,amd64,.exe)
	$(call tar,windows,amd64,.exe)

test/windows:
	$(call test,windows,amd64,.exe)

clean:
	cd build && rm -rf *


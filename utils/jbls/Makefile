# https://golang.org/doc/install/source#environment
GOOS := $(shell go env | awk -F= '$$1=="GOOS" {print $$2}' | awk -F '"' '{print $$2}') # 此处 awk 需使用两个 $
GOARCH := $(shell go env | awk -F= '$$1=="GOARCH" {print $$2}' | awk -F '"' '{print $$2}')
OSS = darwin dragonfly freebsd linux netbsd openbsd plan9 solaris windows
PKG =
# ifeq ($(strip $(GOOS)), windows)
# 	GOARCH := $(strip $(GOARCH)).exe
# endif

.PHONY: build
build:
	@ go build -ldflags "-s -w" -o JetBrainsLicenseServer_$(strip $(GOOS))_$(strip $(if \
    $(findstring windows,$(GOOS)),\
    $(strip $(GOARCH)).exe,\
    $(strip $(GOARCH))\
	))

.PHONY: amd64
amd64:
	@ $(foreach OS,\
	$(OSS),\
	$(shell CGO_ENABLED=0 GOOS=$(OS) GOARCH=amd64 go build -ldflags "-s -w" -o JetBrainsLicenseServer_$(OS)_amd64$(if $(findstring windows,$(OS)),.exe)))
	@ echo done

.PHONY: 386
386:
	@ $(foreach OS,\
	$(OSS),\
	$(shell CGO_ENABLED=0 GOOS=$(OS) GOARCH=386 go build -ldflags "-s -w" -o JetBrainsLicenseServer_$(OS)_386$(if $(findstring windows,$(OS)),.exe)))
	@ echo done

.PHONY: clean
clean:
	@ rm -rf JetBrainsLicenseServer* idea

# 压缩。需要安装 https://github.com/upx/upx
.PHONY: upx
upx:
	@ upx $(if $(PKG),$(PKG),JetBrainsLicenseServer*)
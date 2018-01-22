GOOS := $(shell go env | awk -F= '$$1=="GOOS" {print $$2}' | awk -F '"' '{print $$2}') # 此处 awk 需使用两个 $
GOARCH := $(shell go env | awk -F= '$$1=="GOARCH" {print $$2}' | awk -F '"' '{print $$2}')

.PHONY: build
build:
	@go build -ldflags "-s -w" -o JetBrainsLicenseServer_$(GOOS)_$(GOARCH)
	@upx JetBrainsLicenseServer_$(GOOS)_$(GOARCH)

.PHONY: amd64
amd64:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -o JetBrainsLicenseServer_darwin_amd64
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -o JetBrainsLicenseServer_linux_amd64
	@CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -o JetBrainsLicenseServer_windows_amd64.exe

.PHONY: 386
386:
	@CGO_ENABLED=0 GOOS=darwin GOARCH=386 go build -ldflags "-s -w" -o JetBrainsLicenseServer_darwin_386
	@CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -ldflags "-s -w" -o JetBrainsLicenseServer_linux_386
	@CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -ldflags "-s -w" -o JetBrainsLicenseServer_windows_386.exe

.PHONY: clean
clean:
	@rm -rf JetBrainsLicenseServer* idea

# 压缩。需要安装 https://github.com/upx/upx
.PHONY: upx
upx:
	@upx JetBrainsLicenseServer*
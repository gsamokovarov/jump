NAME = jump
DESCRIPTION = "Jump helps you navigate faster by learning your habits."
HOMEPAGE = https://github.com/gsamokovarov/jump
AUTHOR = "Genadi Samokovarov"
LICENSE = MIT

VERSION = 0.67.0

.PHONY: build
build:
	@env CGO_ENABLED=0 go build -o jump

.PHONY: build.linux
build.linux:
	@env GOOS=linux go build -buildvcs=false -o jump

.PHONY: build.linux.arm
build.linux.arm:
	@env GOOS=linux GOARCH=arm64 go build -o jump

.PHONY: build.arm
build.arm:
	@env GOARCH=arm64 go build -o jump

.PHONY: build.windows
build.windows:
	@env GOOS=windows GOARCH=amd64 go build -o jump.exe

.PHONY: build.osx
build.osx:
	@env GOOS=darwin GOARCH=amd64 go build -o jump

.PHONY: build.osx.arm
build.osx.arm:
	@env GOOS=darwin GOARCH=arm64 go build -o jump

.PHONY: test
test:
	@rm -rf ./config/testdata/.tmp*
	@go test ./... -cover

.PHONY: lint
lint:
	@go vet ./... && golint ./...

.PHONY: clean
clean:
	@rm -f jump*

.PHONY: pkg
pkg: pkg.deb pkg.rpm pkg.linux pkg.linux.arm pkg.windows pkg.osx pkg.osx.arm

.PHONY: pkg.deb
pkg.deb: man build.linux
	@env VERSION=$(VERSION) ARCH=amd64 nfpm package --config nfpm.yaml --packager deb --target .

.PHONY: pkg.deb.arm
pkg.deb.arm: man build.linux.arm
	@env VERSION=$(VERSION) ARCH=arm64 nfpm package --config nfpm.yaml --packager deb --target .

.PHONY: pkg.rpm
pkg.rpm: man build.linux
	@env VERSION=$(VERSION) ARCH=amd64 nfpm package --config nfpm.yaml --packager rpm --target .

.PHONY: pkg.rpm.arm
pkg.rpm.arm: man build.linux.arm
	@env VERSION=$(VERSION) ARCH=arm64 nfpm package --config nfpm.yaml --packager rpm --target .

.PHONY: pkg.linux
pkg.linux: build.linux
	@mv jump jump_linux_amd64_binary

.PHONY: pkg.linux.arm
pkg.linux.arm: man build.linux.arm
	@mv jump jump_linux_arm_binary

.PHONY: pkg.windows
pkg.windows: man build.windows
	@mv jump.exe jump_windows_amd64_binary.exe

.PHONY: pkg.osx
pkg.osx: build.osx
	@mv jump jump_osx

.PHONY: pkg.osx.arm
pkg.osx.arm: build.osx.arm
	@mv jump jump_osx_arm64

.PHONY: man
man: ronn
	@ronn ./man/jump.1.ronn --style=dark
	@cp ./man/jump.1 ./man/j.1

.PHONY: ronn
ronn:
	@which ronn > /dev/null || gem install ronn-ng

NAME = jump
DESCRIPTION = "Jump helps you navigate your file system faster by learning your habits."
HOMEPAGE = https://github.com/gsamokovarov/jump
AUTHOR = "Genadi Samokovarov"
LICENSE = MIT

VERSION = 0.13.0

.PHONY: build
build:
	@go build -o jump

.PHONY: test
test:
	@go test ./... -cover

.PHONY: lint
lint:
	@go vet ./... && golint ./...

# Package deb and rpm inside of a Linux virtual machine, because of
# the user.Current() usage we have. It doesn't work cross-compiled
# from OSX.

.PHONY: deb
deb: build
	@fpm -s dir -t deb -n $(NAME) -v $(VERSION) -a amd64 \
		--deb-compression bzip2 \
		--url $(HOMEPAGE) \
		--description $(DESCRIPTION) \
		--vendor $(AUTHOR) \
		--license $(LICENSE) \
		-m "Genadi Samokovarov <gsamokovarov@gmail.com>" \
		./jump=/usr/bin/jump \
		./man/jump.1=/usr/share/man/man1/jump.1

.PHONY: rpm
rpm: build
	@fpm -s dir -t rpm -n $(NAME) -v $(VERSION) -a amd64 \
		--rpm-compression bzip2 \
		--url $(HOMEPAGE) \
		--description $(DESCRIPTION) \
		--vendor $(AUTHOR) \
		--license $(LICENSE) \
		-m "Genadi Samokovarov <gsamokovarov@gmail.com>" \
		./jump=/usr/bin/jump \
		./man/jump.1=/usr/share/man/man1/jump.1

.PHONY: clean
clean:
	@rm -f jump*

.PHONY: man
man:
	@ronn ./man/jump.1.ronn

NAME = jump
DESCRITPTION = "A fuzzy quick-directory jumper."
HOMEPAGE = https://github.com/gsamokovarov/jump
AUTHOR = "Genadi Samokovarov"
LICENSE = MIT

VERSION = 0.1.0

build:
	@go build -o jump

# Package dep and rpm inside of a Linux virtual machine, because of the
# user.Current() usage we have. It doesn't work cross-compiled from OSX.
deb: build
	@fpm -s dir -t deb -n $(NAME) -v $(VERSION) -a amd64 \
		--deb-compression bzip2 \
		--url $(HOMEPAGE) \
		--description $(DESCRITPTION) \
		--vendor $(AUTHOR) \
		--license $(LICENSE) \
		./jump=/usr/bin/jump

rpm: build
	@fpm -s dir -t rpm -n $(NAME) -v $(VERSION) -a amd64 \
		--rpm-compression bzip2 \
		--url $(HOMEPAGE) \
		--description $(DESCRITPTION) \
		--vendor $(AUTHOR) \
		--license $(LICENSE) \
		./jump=/usr/bin/jump

test:
	@go test ./...

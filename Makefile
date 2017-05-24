NAME="archer"
VERSION="0.0.1"
ITERATION="1"
DESCRIPTION="This is probably how you get ants, Archer"
OUT="./pkg"

all: archer deb

clean:
	rm -rf dist
	rm -rf $(OUT)

deps:
	go get golang.org/x/tools/cmd/goimports
	go get golang.org/x/sys/unix
	go get golang.org/x/net/context
	go get github.com/kardianos/osext
	go get github.com/naegelejd/go-acl/os/group
	go get github.com/boltdb/bolt/...
	go get github.com/chuckpreslar/emission
	go get github.com/mitchellh/cli
	go get github.com/mitchellh/colorstring
	go get github.com/hashicorp/hcl
	go get github.com/hashicorp/hil
	go get github.com/hashicorp/go-multierror
	go get github.com/ryanuber/columnize

archer: clean deps
	mkdir -p dist/etc/archer/collection.d
	mkdir -p dist/usr/lib/archer/package.d
	mkdir -p dist/var/lib/archer
	go build -o dist/usr/bin/archer github.com/instana/archer/cmd/archer

deb:
	which fpm || { echo "fpm must be installed and present in PATH for deb build"; exit 0; }
	mkdir $(OUT)
	fpm -t deb -s dir -C dist -n $(NAME) \
		-v $(VERSION) \
		--description '$(DESCRIPTION)' \
		-a amd64 \
		--iteration $(ITERATION) \
		-p $(OUT) \
		--deb-user root \
		--deb-group root \
		.

fmt:
	goimports -w .

.PHONY: deps fmt

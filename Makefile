
all: archer

clean:
	rm -rf dist

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

fmt:
	goimports -w .

.PHONY: deps fmt

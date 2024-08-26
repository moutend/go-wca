RELEASE_DIR := bin
REVISION    := $(shell git rev-parse --verify HEAD | cut -c-6)
VERSION     ?= latest

export GOOS   := windows
export GOARCH := amd64

examples := $(wildcard example/*)
bins     := $(patsubst example/%,$(RELEASE_DIR)/%.exe,$(examples))

all: $(RELEASE_DIR) $(bins)

$(RELEASE_DIR)/%.exe: $(wildcard example/$*)
	go build -o "$@" -ldflags "-X main.revision=$(REVISION) -X main.version=$(VERSION)" github.com/moutend/go-wca/example/$*

$(RELEASE_DIR):
	mkdir $(RELEASE_DIR)

clean:
	rm -rf $(RELEASE_DIR)
.PHONY: clean

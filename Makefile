RELEASE_DIR=bin
REVISION=$(shell git rev-parse --verify HEAD | cut -c-6)

.PHONY: clean build build-amd64-windows $(DIRNAME) all
all: build-windows-amd64

build-windows-amd64:
	@for v in `ls example`; do \
	$(MAKE) build DIRNAME=$$v GOOS=windows GOARCH=amd64; \
	done;

build-windows-386:
	@echo not supported at the moment

build: $(DIRNAME)

$(DIRNAME):
ifndef VERSION
	@echo '[ERROR] $$VERSION must be specified'
	exit 255
endif
	@cd example/$(DIRNAME);\
	go build -ldflags "-X main.revision=$(REVISION) -X main.version=$(VERSION)"
	@mv example/$(DIRNAME)/$(DIRNAME).exe bin/$(DIRNAME)-$(VERSION)-$(GOARCH).exe

clean:
	rm -rf $(RELEASE_DIR)/*

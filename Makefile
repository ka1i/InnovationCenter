NAME    := InnovationCenter
SOURCE  := cmd/${NAME}/main.go
BINARY  := bin/${NAME}

VERSION := $(shell git describe --tags --always --dirty="-dev")

all: install


.PHONY: install
install:        ## install this app.
	@echo "version:${VERSION}"
	@echo "making ${NAME} ..."
	@go install -ldflags "                                 		\
        -installsuffix 'static'                                 \
        -s -w                                                   \
        -X '$(shell go list -m)/pkg/version.VERSION=${VERSION}' \
        "                                                       \
        ./...

.PHONY: help
help:           ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

.PHONY: clean
clean:          ## Clean build cache.
	@rm -rf bin
	@echo "clean [ ok ]"

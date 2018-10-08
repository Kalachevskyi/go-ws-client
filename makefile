# Build binary
TARGETS = bin/ws-client

WORKDIR := $(PWD)

# Expand name of the module.
MODULE_NAME = ws-client

# Execute "build" command by default
default: build
.PHONY: default

# Build binary file
build: $(TARGETS)
.PHONY: build

bin/ws-client: clean
	@ echo "-> Building $@ binary ..."
	@ go build -o $(WORKDIR)/$@
	@ echo "-> Done!"

# Clean binary file
clean:
	@ echo "-> Cleaning up build artifacts ..."
	@ rm -f $(TARGETS)
	@ echo "-> Done!"
.PHONY: clean

lint:
	@ echo "-> Running linters ..."
	@ golangci-lint run --exclude-use-default=false
	@ echo "-> Done!"
.PHONY: lint
# The default target of the Makefile
# This will be run when you type "make" without any arguments
.PHONY: default
default: lint

# The build target
# This will build the Go program
.PHONY: build
build:
	go build -o myprogram

# The lint target
# This will run the linters and formatters
.PHONY: lint
lint:
	gofmt -s -w .
	golangci-lint run
	go vet ./...

vet:
	go vet ./...
	shadow ./...

# The clean target
# This will remove any generated files
.PHONY: clean
clean:
	rm -f myprogram

# The run target
# This will build and run the Go program
.PHONY: run
run: build
	./myprogram
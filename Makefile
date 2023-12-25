GOPATH = $(shell go env GOPATH)
export GOFLAGS = -mod=mod

all: build
lint:
	which golangci-lint >/dev/null 2>&1 || \
		(curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | \
		sh -s -- -b $(go env GOPATH)/bin)
	golangci-lint run --timeout 10m0s --issues-exit-code 0 | tee report.txt

COVEXCL=	(main.go|offsets|server.go|utils)

test: lint
	[ -f $$GOPATH/bin/go-junit-report ] || go install github.com/jstemmer/go-junit-report@latest
	[ -f $$GOPATH/bin/gocov ] || go install github.com/axw/gocov/gocov@latest
	[ -f $$GOPATH/bin/gocov-xml ] || go install github.com/AlekSi/gocov-xml@latest
	go list ./...
	go vet ./...
	go test -v -race -coverprofile=coverage.tmp ./... | tee test.out
	cat test.out | go-junit-report > test.xml
	egrep -v '$(COVEXCL)' coverage.tmp > coverage.out && rm -f coverage.tmp
	gocov convert coverage.out | tee coverage.json | gocov-xml > coverage.xml
test-flake:
	go test -v -test.failfast -test.count 2 ./...

build: test
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s" -o ./app/ ./cmd/port_domain_service.go

set:
	set

clean:
	rm -rf ./app coverage.json coverage.out coverage.xml test.out test.xml report.txt
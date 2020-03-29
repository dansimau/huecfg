src = $(find . -type f -name '*.go')

huecfg: $(src)
	go build

.PHONY: clean
clean:
	rm -f ./huecfg

.PHONY: test
test:
	go test -v -coverprofile .coverage.out ./...
